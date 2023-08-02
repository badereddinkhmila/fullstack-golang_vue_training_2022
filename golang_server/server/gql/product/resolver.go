package product

import "com.fullstack.ecommerce/server/gql/generated"

type ProductQuery struct {
	productRootQuery
}

type ProductMutation struct {
	productRootMutation
}

type ProductNamespace struct{}

func (ProductNamespace) Product() generated.ProductResolver {
	return &productType{}
}
