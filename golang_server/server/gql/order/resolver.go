package order

import (
	itf "com.fullstack.ecommerce/server/gql/generated"
)

type OrderQuery struct {
	orderRootQuery
}

type OrderMutation struct {
	orderRootMutation
}

type OrderNamespace struct{}

func (OrderNamespace) Order() itf.OrderResolver {
	return &orderType{}
}
