package config

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	appOnce   sync.Once
	appConfig AppConf
)

func App() AppConf {
	appOnce.Do(func() {
		Init()
		appConfig = AppConf{}
		appConfig.load()
	})
	return appConfig
}

type AppConf struct {
	Host string
	Port string
	Cors []string
}

func (obj *AppConf) load() {
	obj.Host = viper.GetString("app.host")
	obj.Port = viper.GetString("app.port")
	obj.Cors = viper.GetStringSlice("app.cors")
}
