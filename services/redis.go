package services

import (
	"RyuLdnWebsite/config"
	"context"
	"flag"
	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
	"log"
)

var RedisClient *rejson.Handler

func InitRedis() {
	rh := rejson.NewReJSONHandler()
	flag.Parse()

	client := redis.NewClient(&redis.Options{
		Addr: config.GetEnv("REDIS_URL"),
	})

	rh.SetGoRedisClientWithContext(context.Background(), client)
	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	RedisClient = rh
}
