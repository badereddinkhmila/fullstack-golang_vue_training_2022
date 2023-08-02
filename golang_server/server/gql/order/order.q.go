package order

import (
	"com.fullstack.ecommerce/app/orders/model"
	"context"
)

type orderRootQuery struct{}

func (orderRootQuery) Orders(ctx context.Context) ([]*model.Order, error) {
	return nil, nil
}

func (orderRootQuery) GetOrderByID(ctx context.Context, id int) (*model.Order, error) {
	return nil, nil
}
