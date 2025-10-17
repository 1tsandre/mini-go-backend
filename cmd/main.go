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
	logger.Infof("Starting application...")

	cfg := initConfig()

	db := initPostgres(cfg)
	defer func() {
		if err := db.Close(); err != nil {
			logger.Errorf("Error closing PostgreSQL: %v", err)
		}
	}()

	rdb := initRedis(cfg)
	defer func() {
		if err := rdb.Close(); err != nil {
			logger.Errorf("Error closing Redis: %v", err)
		}
	}()

	logger.Infof("Application started successfully. PostgreSQL and Redis connected.")
}

func initConfig() *config.Config {
	logger.Infof("Loading configuration...")
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	logger.Infof("Configuration loaded successfully.")
	return cfg
}

func initPostgres(cfg *config.Config) *sql.DB {
	logger.Infof("Connecting to PostgreSQL...")
	db, err := pg.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect PostgreSQL: %v", err)
	}
	logger.Infof("Connected to PostgreSQL successfully.")
	return db
}

func initRedis(cfg *config.Config) *redis.Client {
	logger.Infof("Connecting to Redis...")
	rdb, err := rds.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect Redis: %v", err)
	}
	logger.Infof("Connected to Redis successfully.")
	return rdb
}
