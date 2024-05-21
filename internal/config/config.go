package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

const (
	configPath = "configs"
	configName = "config"
)

func InitConfig() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error with .env file, %s", err.Error())
	}

	return viper.ReadInConfig()
}