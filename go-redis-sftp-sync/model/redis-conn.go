package model

import (
	"os"

	"github.com/go-redis/redis"
)

func CreateRedisConnection() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	return redisClient
}
