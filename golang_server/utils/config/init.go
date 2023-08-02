package config

import (
	"errors"
	"log"
	"sync"

	viper "github.com/spf13/viper"
)

var configOnce sync.Once

func Init() {
	configOnce.Do(func() {
		viper.SetDefault("config", "config.json")

		viper.SetConfigName("config")
		viper.AddConfigPath(".")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			if errors.As(err, &viper.ConfigFileNotFoundError{}) {
				log.Print("Config file not found!")
			} else {
				log.Fatal("Error loading config file: ", err)
			}
		}
	})
}
