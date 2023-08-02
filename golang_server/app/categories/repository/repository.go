package repository

import (
	"com.fullstack.ecommerce/app/categories/model"
	"com.fullstack.ecommerce/utils/postgres"
	"context"
	errors "emperror.dev/errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
)

func GetTableName() string {
	return model.Category{}.TableName()
}

func Create(ctx context.Context, category model.Category) (*model.Category, error) {
	statement := postgres.GetDialect().Insert(GetTableName()).Rows(goqu.Record{
		"lable":       category.Lable,
		"description": category.Description,
		"created_by":  category.CreatedBy,
	}).Returning(goqu.T(GetTableName()).All())

	insertSql, args, err := statement.ToSQL()

	if err != nil {
		return nil, err
	}

	var result = model.Category{}

	if err := pgxscan.Get(ctx, postgres.Client(), &result, insertSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error inserting a new record")
	}
	return &result, nil
}

func FindAll(ctx context.Context) ([]*model.Category, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select()
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}

	var result []*model.Category
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		errors.Wrap(err, "Error selecting records")
	}
	return result, nil
}

func Update(ctx context.Context, category model.Category) (*model.Category, error) {
	statement := postgres.GetDialect().Update(GetTableName()).Set(category).
		Where(goqu.C("id").Eq(category.ID)).Returning(goqu.T(GetTableName()).All())
	updateSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}
	
	var result model.Category
	if err := pgxscan.Get(ctx, postgres.Client(), &result, updateSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error updating record")
	}
	return &result, nil
}

func FindById(ctx context.Context, id int) (*model.Category, error) {
	return FindOneBy(ctx, "id", id)
}

func FindByLable(ctx context.Context, lable string) (*model.Category, error) {
	return FindOneBy(ctx, "lable", lable)
}

func FindBy(ctx context.Context, column_name string, column_value interface{}) ([]*model.Category, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select().Where(goqu.I(column_name).Eq(column_value))
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}
	var result []*model.Category
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error getting records")
	}
	return result, nil
}

func FindOneBy(ctx context.Context, column_name string, column_value interface{}) (*model.Category, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select().Where(goqu.I(column_name).Eq(column_value)).Limit(1)
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}

	var result model.Category
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error getting records")
	}
	if &result == nil {
		return nil, errors.New(fmt.Sprintf("Category with %v:%v does not exist", column_name, column_value))
	}
	return &result, nil
}

func DeleteById(ctx context.Context, id int) (bool, error) {
	statement := postgres.GetDialect().Delete(GetTableName()).Where(goqu.C("id").Eq(id)).
		Returning(goqu.C("id"))
	deleteSql, args, err := statement.ToSQL()
	if err != nil {
		return false, err
	}

	var result int
	if err := pgxscan.Get(ctx, postgres.Client(), &result, deleteSql, args...); err != nil {
		return false, errors.Wrap(err, "Error deleting record")
	}
	return true, nil
}
