package redis

import (
	"github.com/afthaab/job-portal/config"
	"github.com/redis/go-redis/v9"
)

func ConnectToRedis(cfg config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password, // no password set
		DB:       int(cfg.Db),  // use default DB
	})
	return rdb
}
