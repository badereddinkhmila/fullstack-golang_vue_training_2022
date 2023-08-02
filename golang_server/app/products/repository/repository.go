package repository

import (
	"com.fullstack.ecommerce/app/products/model"
	"com.fullstack.ecommerce/utils/postgres"
	"context"
	"emperror.dev/errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
)

func GetTableName() string { return model.Product{}.TableName() }

func Create(ctx context.Context, product model.Product) (*model.Product, error) {
	statement := postgres.GetDialect().Insert(GetTableName()).Rows(goqu.Record{
		"name_product":        product.NameProduct,
		"description_product": product.DescriptionProduct,
		"price_product":       product.Price,
		"quantity":            product.Quantity,
		"is_available":        product.IsAvailable,
		"image_product":       product.ProductImage,
		"category_id":         product.CategoryID,
		"created_by":          product.CreatedBy,
	}).Returning(goqu.T(GetTableName()).All())

	insertSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}

	var result = model.Product{}

	if err := pgxscan.Get(ctx, postgres.Client(), &result, insertSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error inserting a new record")
	}
	return &result, nil
}

func FindAll(ctx context.Context) ([]*model.Product, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select()
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}

	var result = []*model.Product{}
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		errors.Wrap(err, "Error selecting records")
	}
	return result, nil
}

func Update(ctx context.Context, product model.Product) (*model.Product, error) {
	statement := postgres.GetDialect().Update(GetTableName()).Set(product).
		Where(goqu.C("id").Eq(product.ID)).Returning(goqu.T(GetTableName()).All())
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}
	var result = model.Product{}
	if err := pgxscan.Get(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error updating record")
	}
	return &result, nil
}

func DeleteById(ctx context.Context, id int) error {
	statement := postgres.GetDialect().Delete(GetTableName()).Where(goqu.C("id").Eq(id))
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		return err
	}
	var result = model.Product{}
	if err := pgxscan.Get(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		errors.Wrap(err, "Error deleting record")
		return err
	}
	return nil
}

func FindById(ctx context.Context, id int) (*model.Product, error) {
	return FindOneBy(ctx, "id", id)
}

func FindBy(ctx context.Context, column_name string, column_value interface{}) (*model.Product, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select().Where(goqu.I(column_name).Eq(column_value))
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}
	var result []model.Product
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error getting records")
	}
	if len(result) == 0 {
		return nil, errors.New(fmt.Sprintf("Product with %v:%v does not exist", column_name, column_value))
	}
	return &result[0], nil
}

func FindOneBy(ctx context.Context, column_name string, column_value interface{}) (*model.Product, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select().Where(goqu.I(column_name).Eq(column_value)).Limit(1)
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}

	var result model.Product
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error getting records")
	}
	if &result == nil {
		return nil, errors.New(fmt.Sprintf("Product with %v:%v does not exist", column_name, column_value))
	}
	return &result, nil
}

func FindAllByIds(ctx context.Context, ids []int) ([]*model.Product, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select().
		Where(goqu.I("id").In(ids))
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}

	result := []*model.Product{}
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error getting records")
	}
	return result, nil
}
