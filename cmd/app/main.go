package main

import (
	"context"
	"github.com/google/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"pastebin"
	"pastebin/internal/repository"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error with config file: %s", err.Error())
	}

	_, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("error with connecting to database: %s", err.Error())
	}

	srv := pastebin.NewServer(nil, "8080")
	if err := srv.Run(); err != nil {
		logger.Errorf("error with running server: %x", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error with shutting down server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error with .env file, %s", err.Error())
	}

	return viper.ReadInConfig()
}