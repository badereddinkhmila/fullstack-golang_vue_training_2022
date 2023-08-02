package repository

import (
	"context"

	"com.fullstack.ecommerce/app/refreshTokens/model"
	"com.fullstack.ecommerce/utils/postgres"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
)

func Create(ctx context.Context, token *model.RefreshToken) (*model.RefreshToken, error) {
	// fmt.Println(token.ExpiresAt)
	statement := goqu.Dialect("postgres").Insert("refresh_tokens").Rows(goqu.Record{
		"token":      token.Token,
		"expires_at": token.ExpiresAt,
		"user_id":    token.UserID,
	}).Returning(goqu.T("refresh_tokens").All())
	insertSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}
	var result = model.RefreshToken{}
	if err := pgxscan.Get(ctx, postgres.Client(), &result, insertSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error creating a new refresh token!")
	}
	return &result, nil
}

func Update(ctx context.Context, token *model.RefreshToken) (*model.RefreshToken, error) {
	statement := goqu.Dialect("postgres").Update("refresh_tokens").Set(token).Returning(
		goqu.T("refresh_tokens").All())
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}
	var result = model.RefreshToken{}
	if err := pgxscan.Get(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error updating refresh token!")
	}
	return &result, nil
}

func Delete(ctx context.Context, token string) error {
	statement := goqu.Dialect("postgres").Delete("refresh_tokens").Where(goqu.C("token").Eq(token))
	deleteSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}
	var result = model.RefreshToken{}
	if err := pgxscan.Get(ctx, postgres.Client(), &result, deleteSql, args...); err != nil {
		return errors.Wrap(err, "Error deleting record")
	}
	return nil
}

func FindByToken(ctx context.Context, token string) (*model.RefreshToken, error) {
	refreshTokens, err := FindBy(ctx, "token", token)
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching refresh tokens!")
	}
	if len(refreshTokens) == 0 {
		return nil, errors.Wrap(err, "Invalid refresh token!")
	}
	return refreshTokens[0], nil
}

func FindBy(ctx context.Context, column_name string, column_value interface{}) ([]*model.RefreshToken, error) {
	statement := goqu.Dialect("postgres").From("refresh_tokens").Select().Where(goqu.I(column_name).Eq(column_value))
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}
	var result = []*model.RefreshToken{}
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error getting records")
	}
	return result, nil
}
