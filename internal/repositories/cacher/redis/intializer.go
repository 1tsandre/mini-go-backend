package redis

import (
	"context"

	"github.com/1tsandre/mini-go-backend/internal/config"
	"github.com/1tsandre/mini-go-backend/pkg/logger"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		logger.Fatalf("Redis connection failed: %v", err)
	}

	logger.Infof("Connected to Redis")
	return client
}
