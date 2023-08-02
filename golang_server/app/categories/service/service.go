package service

import (
	categModel "com.fullstack.ecommerce/app/categories/model"
	categRepo "com.fullstack.ecommerce/app/categories/repository"
	"context"
	"github.com/xorcare/pointer"
)

func Create(ctx context.Context, category categModel.Category) (*categModel.Category, error) {
	createdCategory, err := categRepo.Create(ctx, category)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

func Update(ctx context.Context, category categModel.Category) (*categModel.Category, error) {
	categoryToUpdate, err := categRepo.FindById(ctx, category.ID)
	if err != nil {
		return nil, err
	}
	categoryToUpdate.Lable = category.Lable
	categoryToUpdate.Description = category.Description
	updatedCategory, err := categRepo.Update(ctx, *categoryToUpdate)
	if err != nil {
		return nil, err
	}
	return updatedCategory, nil

}

func FindAll(ctx context.Context) ([]*categModel.Category, error) {
	categories, err := categRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func FindByID(ctx context.Context, id int) (*categModel.Category, error) {
	category, err := categRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func CheckLableAvailability(ctx context.Context, lable string) *bool {
	category, err := categRepo.FindByLable(ctx, lable)
	return pointer.Bool(err != nil || category.Lable == "")
}

func DeletCategory(ctx context.Context, id int) (bool, error) {
	return categRepo.DeleteById(ctx, id)
}
