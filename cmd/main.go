package main

import (
	"github.com/1tsandre/mini-go-backend/internal/config"
	"github.com/1tsandre/mini-go-backend/internal/repositories/cacher/redis"
	"github.com/1tsandre/mini-go-backend/internal/repositories/database/postgres"
)

func main() {
	cfg := config.LoadConfig()

	db := postgres.NewPostgresConnection(cfg)
	defer db.Close()

	rdb := redis.NewRedisClient(cfg)
	defer rdb.Close()
}
