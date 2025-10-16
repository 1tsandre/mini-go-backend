package redis

import (
	dbRepo "github.com/1tsandre/mini-go-backend/internal/repositories/database"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) dbRepo.CacheRepository {
	return &RedisRepository{redisClient: redisClient}
}
