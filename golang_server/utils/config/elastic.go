package config

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	elasticOnce   sync.Once
	elasticConfig ElasticConf
)

func ElasticConfig() ElasticConf {
	elasticOnce.Do(func() {
		elasticConfig = ElasticConf{}
		elasticConfig.load()
	})
	return elasticConfig
}

type ElasticConf struct {
	Addresses []string
	Username  string
	Password  string
}

func (obj *ElasticConf) load() {
	obj.Addresses = viper.GetStringSlice("elastic.hosts")
	obj.Username = viper.GetString("elastic.username")
	obj.Password = viper.GetString("elastic.password")
}
