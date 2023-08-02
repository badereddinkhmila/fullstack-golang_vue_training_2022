package category

import (
	categoryModel "com.fullstack.ecommerce/app/categories/model"
	categoryService "com.fullstack.ecommerce/app/categories/service"
	"com.fullstack.ecommerce/app/users/model"
	"com.fullstack.ecommerce/server/auth"
	"com.fullstack.ecommerce/server/gql/generated"
	"context"
)

type categoryRootMutation struct{}

func (categoryRootMutation) CreateCategory(ctx context.Context, input *generated.CreateCategoryInput) (*categoryModel.Category, error) {
	//TODO implement me
	currentUser := auth.ForContext(ctx)

	if currentUser == nil || currentUser.Role != model.ROLE_SUPERADMIN {
		return nil, auth.Forbidden(ctx)
	}

	category, err := categoryService.Create(ctx, categoryModel.Category{
		Lable:       input.Lable,
		Description: input.Description,
		CreatedBy:   currentUser.Id,
	})
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (categoryRootMutation) UpdateCategory(ctx context.Context, input *generated.UpdateCategoryInput) (*categoryModel.Category, error) {
	if currentUser := auth.ForContext(ctx); currentUser == nil || currentUser.Role != model.ROLE_SUPERADMIN {
		return nil, auth.Forbidden(ctx)
	}

	category, err := categoryService.Update(ctx, categoryModel.Category{
		ID:          input.ID,
		Lable:       input.Lable,
		Description: input.Description,
	})

	if err != nil {
		return nil, err
	}
	return category, nil
}

func (categoryRootMutation) DeleteCategory(ctx context.Context, id int) (bool, error) {
	if currentUser := auth.ForContext(ctx); currentUser == nil || currentUser.Role != model.ROLE_SUPERADMIN {
		return false, auth.Forbidden(ctx)
	}
	return categoryService.DeletCategory(ctx, id)
}
