package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/1tsandre/mini-go-backend/internal/repositories/cacher"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) cacher.CacheRepository {
	return &RedisRepository{
		client: client,
	}
}

func (r *RedisRepository) Set(ctx context.Context, key string, value interface{}) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisRepository) SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisRepository) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *RedisRepository) IsExist(ctx context.Context, key string) (bool, error) {
	res, err := r.client.Exists(ctx, key).Result()
	return res > 0, err
}

func (r *RedisRepository) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return r.client.Expire(ctx, key, ttl).Err()
}

func (r *RedisRepository) HSet(ctx context.Context, key string, values map[string]interface{}) error {
	return r.client.HSet(ctx, key, values).Err()
}

func (r *RedisRepository) HGet(ctx context.Context, key, field string) (string, error) {
	return r.client.HGet(ctx, key, field).Result()
}

func (r *RedisRepository) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, key).Result()
}

func (r *RedisRepository) Close() error {
	return r.client.Close()
}
