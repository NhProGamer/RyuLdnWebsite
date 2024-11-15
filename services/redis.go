package services

import (
	"RyuLdnWebsite/config"
	"context"
	"flag"
	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var RedisClient *rejson.Handler

func InitRedis() {
	var err error
	rh := rejson.NewReJSONHandler()
	flag.Parse()

	client := redis.NewClient(&redis.Options{
		Addr: config.GetEnv("REDIS_URL"),
	})

	rh.SetGoRedisClientWithContext(context.Background(), client)

	for i := 1; i <= 10; i++ {
		log.Printf("Try to connect to redis server (%v/10)", i)
		err = client.Ping(context.Background()).Err()
		if err != nil {
			log.Printf("Failed to connect to Redis: %v", err)
		} else {
			log.Print("Connected to redis server!")
			break
		}
		time.Sleep(10 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	RedisClient = rh
}
