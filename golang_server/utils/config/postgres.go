package config

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	postgresOnce sync.Once
	pgConfig     PostgresConf
)

func PostgreSQL() PostgresConf {
	postgresOnce.Do(func() {
		pgConfig = PostgresConf{}
		pgConfig.load()
	})
	return pgConfig
}

type PostgresConf struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (obj *PostgresConf) load() {
	obj.Host = viper.GetString("db.host")
	obj.Port = viper.GetInt("db.port")
	obj.Username = viper.GetString("db.username")
	obj.Password = viper.GetString("db.password")
	obj.Database = viper.GetString("db.database")
}
