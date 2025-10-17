package app

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/1tsandre/mini-go-backend/internal/config"
	httpserver "github.com/1tsandre/mini-go-backend/internal/handlers/http"
	rds "github.com/1tsandre/mini-go-backend/internal/repositories/cacher/redis"
	pg "github.com/1tsandre/mini-go-backend/internal/repositories/database/postgres"
	"github.com/1tsandre/mini-go-backend/pkg/logger"
)

type App struct {
	cfg    *config.Config
	db     *sql.DB
	rdb    *redis.Client
	server *httpserver.Server
}

func New(cfg *config.Config) (*App, error) {
	db, err := pg.New(cfg)
	if err != nil {
		return nil, err
	}

	rdbClient, err := rds.New(cfg)
	if err != nil {
		return nil, err
	}

	router := httpserver.NewRouter()
	server := httpserver.NewHTTPServer(cfg.Server.Port, router)

	return &App{
		cfg:    cfg,
		db:     db,
		rdb:    rdbClient,
		server: server,
	}, nil
}

func (a *App) Start() {
	go func() {
		logger.Infof("HTTP server starting on port %d", a.cfg.Server.Port)
		if err := a.server.Start(); err != nil {
			logger.Errorf("HTTP server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Infof("Shutting down application...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		logger.Errorf("Error shutting down HTTP server: %v", err)
	}

	if err := a.rdb.Close(); err != nil {
		logger.Errorf("Error closing Redis: %v", err)
	}

	if err := a.db.Close(); err != nil {
		logger.Errorf("Error closing PostgreSQL: %v", err)
	}

	logger.Infof("Shutdown complete.")
}
