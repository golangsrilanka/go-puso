package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error reading config file, %s", err)
	}

	env := viper.GetString(key)

	return env
}
