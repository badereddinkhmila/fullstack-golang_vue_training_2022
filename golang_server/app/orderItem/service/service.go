package service

import (
	"context"

	ordItemModel "com.fullstack.ecommerce/app/orderItem/model"
	ordItemRepo "com.fullstack.ecommerce/app/orderItem/repository"
)

func Create(ctx context.Context, ordItem ordItemModel.OrderItem) (ordItemModel.OrderItem, error) {
	createdOrdItem, err := ordItemRepo.Create(ctx, ordItem)
	if err != nil {
		return ordItemModel.OrderItem{}, err
	}
	return createdOrdItem, nil
}
