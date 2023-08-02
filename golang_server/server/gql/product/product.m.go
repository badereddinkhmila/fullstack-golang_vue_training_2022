package product

import (
	"com.fullstack.ecommerce/app/products/model"
	productModel "com.fullstack.ecommerce/app/products/model"
	productService "com.fullstack.ecommerce/app/products/service"
	userModel "com.fullstack.ecommerce/app/users/model"
	"com.fullstack.ecommerce/server/auth"
	"com.fullstack.ecommerce/server/gql/generated"
	"context"
)

type productRootMutation struct{}

func (productRootMutation) CreateProduct(ctx context.Context, input *generated.CreateProductInput) (*model.Product, error) {
	currentUser := auth.ForContext(ctx)
	if currentUser == nil || (currentUser.Role != userModel.ROLE_SUPERADMIN && currentUser.Role != userModel.ROLE_PROVIDER) {
		return nil, auth.Forbidden(ctx)
	}

	product, err := productService.Create(ctx, productModel.Product{
		NameProduct:        input.NameProduct,
		DescriptionProduct: input.DescriptionProduct,
		Price:              input.PriceProduct,
		Quantity:           input.QuantityProduct,
		IsAvailable:        input.IsAvailable,
		CategoryID:         input.CategoryID,
		ProductImage:       []byte(input.ProductImage),
		CreatedBy:          currentUser.Id,
	})

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (productRootMutation) UpdateProduct(ctx context.Context, input *generated.UpdateProductInput) (*model.Product, error) {
	currentUser := auth.ForContext(ctx)
	if currentUser == nil || (currentUser.Role != userModel.ROLE_SUPERADMIN && currentUser.Role != userModel.ROLE_PROVIDER) {
		return nil, auth.Forbidden(ctx)
	}
	return productService.Update(ctx, model.Product{
		ID:                 input.ID,
		NameProduct:        input.NameProduct,
		DescriptionProduct: input.DescriptionProduct,
		Price:              input.PriceProduct,
		Quantity:           input.QuantityProduct,
		IsAvailable:        input.IsAvailable,
		CategoryID:         input.CategoryID,
		CreatedBy:          input.CreatedBy,
	}, currentUser)
}

func (productRootMutation) DeleteProduct(ctx context.Context, id int) (bool, error) {
	currentUser := auth.ForContext(ctx)
	if currentUser == nil || (currentUser.Role != userModel.ROLE_SUPERADMIN && currentUser.Role != userModel.ROLE_PROVIDER) {
		return false, auth.Forbidden(ctx)
	}
	return productService.DeletProduct(ctx, id, *currentUser)
}
