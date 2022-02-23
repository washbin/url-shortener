package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.TODO()

func Init(address string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
