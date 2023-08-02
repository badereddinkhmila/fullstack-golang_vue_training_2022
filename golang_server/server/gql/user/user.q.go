package user

import (
	"com.fullstack.ecommerce/app/users/model"
	"com.fullstack.ecommerce/app/users/service"
	"com.fullstack.ecommerce/server/auth"
	"com.fullstack.ecommerce/server/gql/generated"
	"context"
)

type userRootQuery struct{}

func (q userRootQuery) Users(ctx context.Context) ([]*model.User, error) {
	//TODO implement me
	if currentUser := auth.ForContext(ctx); currentUser == nil || currentUser.Role != model.ROLE_SUPERADMIN {
		return nil, auth.Forbidden(ctx)
	}

	users, err := service.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (q userRootQuery) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	//TODO implement me
	if currentUser := auth.ForContext(ctx); currentUser == nil || currentUser.Role != model.ROLE_SUPERADMIN {
		return nil, auth.Forbidden(ctx)
	}
	user, err := service.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (q userRootQuery) GetMe(ctx context.Context) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (q userRootQuery) CheckUsernameAvailability(ctx context.Context, username string) (*bool, error) {
	//TODO implement me
	available := service.CheckUsernameAvailability(ctx, username)
	return &available, nil
}

func (q userRootQuery) CheckEmailAvailability(ctx context.Context, email string) (*bool, error) {
	//TODO implement me
	available := service.CheckEmailAvailability(ctx, email)
	return &available, nil
}

func (q userRootQuery) LoginUser(ctx context.Context, input generated.LoginInput) (*model.LoginResponse, error) {
	//TODO implement me
	return service.LoginUser(ctx, model.UserCredentials{
		UsernameOrEmail: input.UsernameOrEmail,
		Password:        input.Password,
	})
}

func (q userRootQuery) RefreshToken(ctx context.Context, token string) (*model.LoginResponse, error) {
	//TODO implement me
	return service.RefreshUserTokens(ctx, token)
}
