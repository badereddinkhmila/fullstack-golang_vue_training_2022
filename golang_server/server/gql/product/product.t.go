package product

import (
	productModel "com.fullstack.ecommerce/app/products/model"
	userModel "com.fullstack.ecommerce/app/users/model"
	userService "com.fullstack.ecommerce/app/users/service"
	"context"
)

type productType struct{}

func (t productType) PriceProduct(ctx context.Context, obj *productModel.Product) (float64, error) {
	//TODO implement me
	return obj.Price, nil
}

func (t productType) QuantityProduct(ctx context.Context, obj *productModel.Product) (int, error) {
	//TODO implement me
	return obj.Quantity, nil
}

func (t productType) ProductImage(ctx context.Context, obj *productModel.Product) (string, error) {
	//TODO implement me
	return string(obj.ProductImage), nil
}

func (productType) CreatedBy(ctx context.Context, obj *productModel.Product) (*userModel.User, error) {
	user, err := userService.GetUserById(ctx, obj.CreatedBy)
	if err != nil {
		return nil, err
	}
	return user, nil
}
