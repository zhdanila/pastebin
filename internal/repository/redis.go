package repository

import (
	"context"
	"encoding/base64"
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

func GenerateHash(client *redis.Client, amount int) {
	for i := 1; i <= amount; i++ {
		hash := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", i)))

		err := client.HSet(context.Background(), "myHash", fmt.Sprintf("%d", i), hash).Err()
		if err != nil {
			fmt.Println("Error writing to Redis:", err)
		}
	}
}