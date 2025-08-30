package utils

import "github.com/redis/go-redis/v9"

var redisClient *redis.Client

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
}

func GetRedisInstance() *redis.Client {
	return redisClient
}
