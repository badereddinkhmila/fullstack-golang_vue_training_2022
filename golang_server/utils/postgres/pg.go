package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"com.fullstack.ecommerce/utils/config"
)

var (
	pgxSyncOnce sync.Once
	pgxp        *pgxpool.Pool
)

func connect() {
	c := config.PostgreSQL()
	config := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		c.Username, c.Password, c.Host, c.Port, c.Database, "disable",
	)

	pgxConfig, err := pgxpool.ParseConfig(config)
	if err != nil {
		log.Fatalf("Error parsing pgx config: %v", err)
	}

	connection, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatalf("Unable to connect PGX: %v", err)
	}
	pgxp = connection
}

func Client() *pgxpool.Pool {
	pgxSyncOnce.Do(func() {
		connect()
	})
	return pgxp
}

func RecordNoFound(err error) bool {
	return pgxscan.NotFound(err) || errors.Is(err, pgx.ErrNoRows)
}

func GetDialect() goqu.DialectWrapper { return goqu.Dialect("postgres") }
