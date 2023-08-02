package repository

import (
	"com.fullstack.ecommerce/app/users/model"
	"com.fullstack.ecommerce/utils/postgres"
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
	"regexp"
)

func GetTableName() string { return model.User{}.TableName() }

func Create(ctx context.Context, user model.User) (model.User, error) {
	statement := postgres.GetDialect().Insert(GetTableName()).Rows(goqu.Record{
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
		"role":     user.Role,
	}).Returning(goqu.T(GetTableName()).All())

	insertSql, args, err := statement.ToSQL()

	if err != nil {
		return model.User{}, err
	}

	var result = model.User{}

	if err := pgxscan.Get(ctx, postgres.Client(), &result, insertSql, args...); err != nil {
		return result, errors.Wrap(err, "Error inserting a new record")
	}

	return result, nil
}

func FindAll(ctx context.Context) ([]*model.User, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select()
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}

	var result = []*model.User{}
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error selecting records")
	}
	return result, nil
}

func FindById(ctx context.Context, id int) (*model.User, error) {
	return FindOneBy(ctx, "id", id)
}

func FindByUsername(ctx context.Context, username string) (*model.User, error) {
	return FindOneBy(ctx, "username", username)
}

func FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return FindOneBy(ctx, "email", email)
}
func FindByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*model.User, error) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(usernameOrEmail) {
		return FindByEmail(ctx, usernameOrEmail)
	}
	return FindByUsername(ctx, usernameOrEmail)
}

func FindOneBy(ctx context.Context, column_name string, column_value interface{}) (*model.User, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select().Where(goqu.I(column_name).Eq(column_value)).Limit(1)
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		return nil, err
	}
	result := model.User{}
	if err := pgxscan.Get(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error getting records")
	}
	if &result == nil {
		return nil, errors.New(fmt.Sprintf("User with %v:%v does not exist", column_name, column_value))
	}
	return &result, nil
}

func Update(ctx context.Context, user model.User) (*model.User, error) {
	statement := postgres.GetDialect().Update(GetTableName()).Set(user).
		Where(goqu.C("id").Eq(user.Id)).Returning(goqu.T(GetTableName()).All())
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}
	var result = model.User{}
	if err := pgxscan.Get(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error updating record")
	}
	return &result, nil
}

func DeleteById(ctx context.Context, id int) error {
	statement := postgres.GetDialect().Delete(GetTableName()).Where(goqu.C("id").Eq(id))
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}
	var result = model.User{}
	if err := pgxscan.Get(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return errors.Wrap(err, "Error deleting record")
	}
	return nil
}

func FindBy(ctx context.Context, column_name string, column_value interface{}) ([]*model.User, error) {
	statement := postgres.GetDialect().From(GetTableName()).Select().Where(goqu.I(column_name).Eq(column_value))
	selectSql, args, err := statement.ToSQL()
	if err != nil {
		panic(err)
	}
	var result []*model.User
	if err := pgxscan.Select(ctx, postgres.Client(), &result, selectSql, args...); err != nil {
		return nil, errors.Wrap(err, "Error getting records")
	}
	return result, nil
}
