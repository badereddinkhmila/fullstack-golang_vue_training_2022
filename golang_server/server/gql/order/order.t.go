package order

import (
	"context"

	orderModel "com.fullstack.ecommerce/app/orders/model"
	userModel "com.fullstack.ecommerce/app/users/model"
)

type orderType struct{}

func (t orderType) Status(ctx context.Context, obj *orderModel.Order) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (orderType) CreatedBy(ctx context.Context, obj *orderModel.Order) (*userModel.User, error) {
	return nil, nil
}
