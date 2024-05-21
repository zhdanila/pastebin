package app

import (
	"context"
	"github.com/google/logger"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"pastebin"
	"pastebin/internal/config"
	"pastebin/internal/repository"
	"syscall"
)

func Run() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error with config file: %s", err.Error())
	}

	//postgres
	_, err := repository.NewPostgresDB(repository.PostgresConfig{
		Host:     viper.GetString("postgresql_db.host"),
		Port:     viper.GetString("postgresql_db.port"),
		Username: viper.GetString("postgresql_db.username"),
		DBName:   viper.GetString("postgresql_db.dbname"),
		SSLMode:  viper.GetString("postgresql_db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("error with connecting to postresql database: %s", err.Error())
	}

	//redis
	_, err = repository.NewRedisDB(repository.RedisConfig{
		Port: viper.GetString("redis_db.port"),
	})
	if err != nil {
		log.Fatalf("error with connecting to redis database: %s", err.Error())
	}

	//amazons3
	_, err = repository.NewAmazonDB(repository.AmazonConfig{
		Region:    viper.GetString("amazon_db.region"),
		AccessKey: viper.GetString("amazon_db.access-key"),
		SecretKey: viper.GetString("amazon_db.secret-access-key"),
	})
	if err != nil {
		log.Fatalf("error with connecting to amazons3 database: %s", err.Error())
	}

	//starting server
	srv := pastebin.NewServer(nil, "8080")
	if err := srv.Run(); err != nil {
		logger.Errorf("error with running server: %x", err)
	}

	//graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error with shutting down server: %s", err.Error())
	}
}
