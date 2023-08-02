package service

import (
	"context"

	ordModel "com.fullstack.ecommerce/app/orders/model"
	ordRepo "com.fullstack.ecommerce/app/orders/repository"
)

func Create(ctx context.Context, order ordModel.Order) (ordModel.Order, error) {
	createdOrd, err := ordRepo.Create(ctx, order)
	if err != nil {
		return ordModel.Order{}, err
	}
	return createdOrd, nil
}
