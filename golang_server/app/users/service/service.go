package service

import (
	"context"
	"time"

	"github.com/spf13/viper"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	tokensModel "com.fullstack.ecommerce/app/refreshTokens/model"
	tokensService "com.fullstack.ecommerce/app/refreshTokens/service"
	userModel "com.fullstack.ecommerce/app/users/model"
	userRepo "com.fullstack.ecommerce/app/users/repository"
	"com.fullstack.ecommerce/utils/helpers"
)

func Create(ctx context.Context, user userModel.User) (userModel.User, error) {
	user.Role = userModel.ROLE_CUSTOMER
	createdUser, err := userRepo.Create(ctx, user)
	if err != nil {
		return userModel.User{}, err
	}
	return createdUser, nil
}
func CreateProvider(ctx context.Context, user userModel.User) (userModel.User, error) {
	user.Role = userModel.ROLE_PROVIDER
	createdProvider, err := userRepo.Create(ctx, user)
	if err != nil {
		return userModel.User{}, err
	}
	return createdProvider, nil
}
func LoginUser(ctx context.Context, credentials userModel.UserCredentials) (*userModel.LoginResponse, error) {
	user, err := userRepo.FindByUsernameOrEmail(ctx, credentials.UsernameOrEmail)

	if err != nil {
		return nil, errors.Wrap(err, "Invalid credentials!")
	}
	if user.Username == "" {
		return nil, errors.New("Invalid credentials!")
	}
	if !helpers.CheckPasswordHash(credentials.Password, user.Password) {
		return nil, errors.New("Invalid user credentials!")
	}
	accesToken, err := helpers.GenerateAccessToken(user.Username)
	if err != nil {
		return nil, err
	}

	refreshToken, err := tokensService.Create(ctx, tokensModel.RefreshToken{
		Token:     uuid.New().String(),
		ExpiresAt: time.Now().Unix() + viper.GetInt64("jwt.refreshTokenExpiry"),
		UserID:    user.Id,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error creating refresh token")
	}

	return &userModel.LoginResponse{AccessToken: accesToken, RefreshToken: refreshToken.Token, CurrentUser: *user}, nil
}

func RefreshUserTokens(ctx context.Context, refreshToken string) (*userModel.LoginResponse, error) {
	token, err := tokensService.GetByToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	if token.ExpiresAt < time.Now().Unix() {
		// check err
		tokensService.Delete(ctx, refreshToken)
		return nil, errors.New("This refresh token is expired !")
	}
	user, err := GetUserById(ctx, token.UserID)

	if err != nil {
		return nil, err
	}
	accesToken, err := helpers.GenerateAccessToken(user.Username)
	if err != nil {
		return nil, err
	}
	return &userModel.LoginResponse{AccessToken: accesToken, RefreshToken: refreshToken, CurrentUser: *user}, nil
}

func FindAll(ctx context.Context) ([]*userModel.User, error) {
	users, err := userRepo.FindAll(ctx)
	if err != nil {
		return users, err
	}
	return users, nil
}

func GetUserById(ctx context.Context, id int) (*userModel.User, error) {

	user, err := userRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CheckUsernameAvailability(ctx context.Context, username string) bool {
	user, err := userRepo.FindByUsername(ctx, username)
	if err != nil || user.Username == "" {
		return true
	}
	return false
}

func CheckEmailAvailability(ctx context.Context, email string) bool {
	user, err := userRepo.FindByEmail(ctx, email)
	/*
		if err != nil || user.Email == "" {
			return true
		}
		return false*/

	return err != nil || user.Email == ""
}
