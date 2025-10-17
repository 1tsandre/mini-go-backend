package main

import (
	"database/sql"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/1tsandre/mini-go-backend/internal/config"
	rds "github.com/1tsandre/mini-go-backend/internal/repositories/cacher/redis"
	pg "github.com/1tsandre/mini-go-backend/internal/repositories/database/postgres"
	"github.com/1tsandre/mini-go-backend/pkg/logger"
)

func main() {
	cfg := initConfig()

	db := initPostgres(cfg)
	defer db.Close()

	rdb := initRedis(cfg)
	defer rdb.Close()

	logger.Infof("Application started successfully. PostgreSQL and Redis connected.")
}

func initConfig() *config.Config {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	return cfg
}

func initPostgres(cfg *config.Config) *sql.DB {
	db, err := pg.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect PostgreSQL: %v", err)
	}
	return db
}

func initRedis(cfg *config.Config) *redis.Client {
	rdb, err := rds.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect Redis: %v", err)
	}
	return rdb
}
