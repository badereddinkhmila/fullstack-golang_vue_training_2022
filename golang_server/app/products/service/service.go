package service

import (
	prodModel "com.fullstack.ecommerce/app/products/model"
	prodRepo "com.fullstack.ecommerce/app/products/repository"
	userModel "com.fullstack.ecommerce/app/users/model"
	"com.fullstack.ecommerce/server/auth"
	"context"
)

func Create(ctx context.Context, product prodModel.Product) (*prodModel.Product, error) {
	createdProduct, err := prodRepo.Create(ctx, product)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
}

func GetAllProducts(ctx context.Context) ([]*prodModel.Product, error) {
	products, err := prodRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductById(ctx context.Context, id int) (*prodModel.Product, error) {
	product, err := prodRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func Update(ctx context.Context, product prodModel.Product, currentUser *userModel.User) (*prodModel.Product, error) {
	if currentUser == nil || currentUser.Id != product.CreatedBy {

	}
	updatedProduct, err := prodRepo.Update(ctx, product)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func DeletProduct(ctx context.Context, id int, currentUser userModel.User) (bool, error) {
	if product, err := GetProductById(ctx, id); err == nil || product.CreatedBy != currentUser.Id {
		return false, auth.Forbidden(ctx)
	}
	err := prodRepo.DeleteById(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
