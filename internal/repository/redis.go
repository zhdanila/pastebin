package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisConfig struct {
	Port string
}

func NewRedisDB(config RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%s", config.Port),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}