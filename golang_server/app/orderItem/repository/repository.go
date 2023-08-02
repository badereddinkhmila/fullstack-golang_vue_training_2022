package repository

import (
	"context"

	"com.fullstack.ecommerce/app/orderItem/model"
	"com.fullstack.ecommerce/utils/postgres"
	errors "emperror.dev/errors"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
)

func GetTableName() string {
	return model.OrderItem{}.TableName()
}

func Create(ctx context.Context, orderItem model.OrderItem) (model.OrderItem, error) {
	statement := postgres.GetDialect().Insert(GetTableName()).Rows(goqu.Record{
		"OrderID":   orderItem.OrderID,
		"Quantity":  orderItem.Quantity,
		"ProductID": orderItem.ProductID,
	}).Returning(goqu.T(GetTableName()).All())

	insertSql, args, err := statement.ToSQL()

	if err != nil {
		panic(err)
	}

	var result = model.OrderItem{}

	if err := pgxscan.Get(ctx, postgres.Client(), &result, insertSql, args...); err != nil {
		errors.Wrap(err, "Error inserting a new record")
	}
	return result, nil
}

func FindAll(ctx context.Context) ([]*model.OrderItem, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select()
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}

	var result = []*model.OrderItem{}
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		errors.Wrap(err, "Error selecting records")
	}
	return result, nil
}

func Update(ctx context.Context, orderItem model.OrderItem) (*model.OrderItem, error) {
	statement := postgres.GetDialect().Update(GetTableName()).Set(orderItem)
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}
	var result = model.OrderItem{}
	if err := pgxscan.Get(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		errors.Wrap(err, "Error updating record")
	}
	return &result, nil
}

func DeleteById(ctx context.Context, id int) error {
	statement := postgres.GetDialect().Delete(GetTableName()).Where(goqu.C("id").Eq(id))
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}
	var result = model.OrderItem{}
	if err := pgxscan.Get(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		errors.Wrap(err, "Error deleting record")
		return err
	}
	return nil
}
