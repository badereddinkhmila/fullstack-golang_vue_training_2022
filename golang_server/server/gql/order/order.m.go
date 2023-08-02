package order

import (
	orderModel "com.fullstack.ecommerce/app/orders/model"
	"com.fullstack.ecommerce/server/gql/generated"
	"context"
)

type orderRootMutation struct{}

func (orderRootMutation) CreateOrder(ctx context.Context, input *generated.CreateOrderInput) (*orderModel.Order, error) {
	//TODO implement me
	return nil, nil
}

func (orderRootMutation) UpdateOrder(ctx context.Context, input *generated.UpdateOrderInput) (*orderModel.Order, error) {
	return nil, nil
}

func (orderRootMutation) DeleteOrder(ctx context.Context, id int) (bool, error) {
	return false, nil
}
