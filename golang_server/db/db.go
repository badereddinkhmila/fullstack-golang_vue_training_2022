package db

import (
	"context"
	"embed"
	"path"
	"strings"

	"com.fullstack.ecommerce/utils/postgres"
)

//go:embed schema/*
var sqlFs embed.FS

var sqlDir = "schema"

func Init(ctx context.Context) error {
	files, err := sqlFs.ReadDir(sqlDir)
	if err != nil {
		panic(err)
	}

	runSQL := func(suffix string) error {
		for _, file := range files {
			if strings.Contains(file.Name(), suffix) {
				if err = RunSqlSchemaFile(ctx, file.Name()); err != nil {
					return err
				}
			}
		}
		return nil
	}

	if err := runSQL("table"); err != nil {
		return err
	} else if err := runSQL("alter"); err != nil {
		return err
	}
	return nil
}

func RunSqlSchemaFile(ctx context.Context, filename string) error {
	filePath := path.Join(sqlDir, filename)
	if sql, err := sqlFs.ReadFile(filePath); err != nil {
		return err
	} else if _, err = postgres.Client().Exec(ctx, string(sql)); err != nil {
		return err
	}
	return nil
}
