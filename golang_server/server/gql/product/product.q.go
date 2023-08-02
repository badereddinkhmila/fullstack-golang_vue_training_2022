package product

import (
	"com.fullstack.ecommerce/app/products/model"
	productService "com.fullstack.ecommerce/app/products/service"
	"context"
)

type productRootQuery struct{}

func (productRootQuery) Products(ctx context.Context) ([]*model.Product, error) {
	return productService.GetAllProducts(ctx)
}

func (productRootQuery) GetProductByID(ctx context.Context, id int) (*model.Product, error) {
	return productService.GetProductById(ctx, id)
}
