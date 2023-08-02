package category

import (
	categoryModel "com.fullstack.ecommerce/app/categories/model"
	userModel "com.fullstack.ecommerce/app/users/model"
	userService "com.fullstack.ecommerce/app/users/service"
	"context"
)

type categoryType struct{}

func (categoryType) CreatedBy(ctx context.Context, obj *categoryModel.Category) (*userModel.User, error) {
	user, err := userService.GetUserById(ctx, obj.CreatedBy)
	if err != nil {
		return nil, err
	}
	return user, nil
}
