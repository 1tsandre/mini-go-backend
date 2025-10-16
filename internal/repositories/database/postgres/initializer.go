package postgres

import (
	"database/sql"
	"fmt"

	"github.com/1tsandre/mini-go-backend/internal/config"
	"github.com/1tsandre/mini-go-backend/pkg/logger"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatalf("Failed to open PostgreSQL: %v", err)
	}

	if err := db.Ping(); err != nil {
		logger.Fatalf("Cannot ping PostgreSQL: %v", err)
	}

	logger.Infof("Connected to PostgreSQL")
	return db
}
