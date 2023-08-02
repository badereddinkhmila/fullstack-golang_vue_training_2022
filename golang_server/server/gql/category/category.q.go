package category

import (
	"com.fullstack.ecommerce/app/categories/model"
	categoryService "com.fullstack.ecommerce/app/categories/service"
	"context"
)

type categoryRootQuery struct{}

func (categoryRootQuery) Categories(ctx context.Context) ([]*model.Category, error) {
	return categoryService.FindAll(ctx)
}
func (categoryRootQuery) GetCategoryByID(ctx context.Context, id int) (*model.Category, error) {
	return categoryService.FindByID(ctx, id)
}
func (categoryRootQuery) CheckLableAvailability(ctx context.Context, lable string) (*bool, error) {
	return categoryService.CheckLableAvailability(ctx, lable), nil
}
