package cache

import (
	"context"

	"github.com/lukanzx/DouServer/config"
	"github.com/lukanzx/DouServer/pkg/constants"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       constants.RedisDBVideo,
	})
	_, err := RedisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
}
