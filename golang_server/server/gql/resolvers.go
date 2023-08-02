package gql

import (
	"com.fullstack.ecommerce/server/gql/category"
	generated "com.fullstack.ecommerce/server/gql/generated"
	"com.fullstack.ecommerce/server/gql/order"
	"com.fullstack.ecommerce/server/gql/orderitem"
	"com.fullstack.ecommerce/server/gql/product"
	"com.fullstack.ecommerce/server/gql/user"
)

type resolver struct{}

func Resolver() generated.ResolverRoot {
	return &struct {
		resolver
		category.CategoryNamespace
		product.ProductNamespace
		order.OrderNamespace
	}{}
}

func (resolver) Mutation() generated.MutationResolver {
	return &struct {
		user.UserMutation
		category.CategoryMutation
		product.ProductMutation
		orderitem.OrderItemMutation
		order.OrderMutation
	}{}
}

func (resolver) Query() generated.QueryResolver {
	return &struct {
		user.UserQuery
		category.CategoryQuery
		product.ProductQuery
		orderitem.OrderItemQuery
		order.OrderQuery
	}{}
}
