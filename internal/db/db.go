package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db     *sql.DB
	logger *slog.Logger
}

const (
	DATABASE_CONNECTION_TRIES_COUNT = 10
)

func GetStorage(logger slog.Logger) (Storage, error) {
	fn := "internal.db.db.GetStorage"
	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	postgres_db := os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf("postgres://%s:%s@postgres:5432/%s", postgres_user, postgres_password, postgres_db)

	var db *sql.DB
	var err error

	for i := range DATABASE_CONNECTION_TRIES_COUNT {
		// Connection with data base
		db, err = sql.Open("pgx", dsn)

		if err != nil {
			logger.Warn(fmt.Sprintf("Cannot open database, try number: %d/%d", i+1, DATABASE_CONNECTION_TRIES_COUNT))
			continue
		}

		if err := db.Ping(); err != nil {
			logger.Warn(fmt.Sprintf("Cannot ping database, try number: %d/%d", i+1, DATABASE_CONNECTION_TRIES_COUNT))
			continue
		}
		logger.Info("esablished connection with database")
		break
	}

	if err != nil {
		return Storage{}, fmt.Errorf("%s:%w", fn, err)
	}

	storage := Storage{
		logger: logger.With("service", "DataBase"),
		db:     db,
	}

	// tasks table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		text TEXT NOT NULL
	);
	`)

	if err != nil {
		return Storage{}, fmt.Errorf("cannot create notes table: %s", err.Error())
	}

	return storage, nil
}
