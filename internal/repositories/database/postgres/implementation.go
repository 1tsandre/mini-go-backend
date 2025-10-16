package postgres

import (
	"database/sql"

	dbRepo "github.com/1tsandre/mini-go-backend/internal/repositories/database"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) dbRepo.DatabaseRepository {
	return &PostgresRepository{db: db}
}