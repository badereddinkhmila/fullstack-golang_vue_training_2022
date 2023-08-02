package service

import (
	"context"

	"com.fullstack.ecommerce/app/refreshTokens/model"
	tokensRepo "com.fullstack.ecommerce/app/refreshTokens/repository"
)

func Create(ctx context.Context, token model.RefreshToken) (*model.RefreshToken, error) {
	createdToken, err := tokensRepo.Create(ctx, &token)
	if err != nil {
		return nil, err
	}
	return createdToken, nil
}

func GetByToken(ctx context.Context, token string) (*model.RefreshToken, error) {
	resultToken, err := tokensRepo.FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return resultToken, nil
}

func Delete(ctx context.Context, token string) error {
	err := tokensRepo.Delete(ctx, token)
	if err != nil {
		return err
	}
	return nil
}
