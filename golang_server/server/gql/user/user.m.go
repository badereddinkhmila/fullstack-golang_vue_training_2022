package user

import (
	"com.fullstack.ecommerce/server/auth"
	"context"
	"fmt"

	"com.fullstack.ecommerce/server/gql/generated"
	"com.fullstack.ecommerce/utils/helpers"

	"com.fullstack.ecommerce/app/users/model"
	userService "com.fullstack.ecommerce/app/users/service"
)

type userRootMutation struct{}

func (m userRootMutation) RegisterCustomer(ctx context.Context, input *generated.CreateUserInput) (*model.User, error) {
	//TODO implement me
	hash, err := helpers.HashPassword(input.Password)
	if err != nil {
		panic(err)
	}

	user, err := userService.Create(ctx, model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hash,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m userRootMutation) RegisterProvider(ctx context.Context, input *generated.CreateUserInput) (*model.User, error) {
	//TODO implement me
	if currentUser := auth.ForContext(ctx); currentUser == nil || currentUser.Role != model.ROLE_SUPERADMIN {
		return nil, auth.Forbidden(ctx)
	}
	hash, err := helpers.HashPassword(input.Password)
	if err != nil {
		panic(err)
	}

	user, err := userService.Create(ctx, model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hash,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m userRootMutation) UpdateUser(ctx context.Context, input *generated.UpdateUserInput) (*model.User, error) {
	//TODO implement me
	if currentUser := auth.ForContext(ctx); currentUser == nil {
		return nil, auth.Forbidden(ctx)
	}
	fmt.Println("from update user")
	return nil, nil
}

func (m userRootMutation) ResetPassword(ctx context.Context, input *generated.ResetPasswordInput) (*model.User, error) {
	//TODO implement me
	if currentUser := auth.ForContext(ctx); currentUser == nil {
		return nil, auth.Forbidden(ctx)
	}
	fmt.Println("from reset password")
	return nil, nil
}

func (m userRootMutation) DeleteUser(ctx context.Context, id int) (bool, error) {
	fmt.Println("from delete user")
	return false, nil
}
