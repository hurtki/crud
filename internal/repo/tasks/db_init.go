package tasks_repo

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type TaskStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

const (
	DATABASE_CONNECTION_TRIES_COUNT = 10
	RETRY_TIME_BETWEEN_TRIES        = 1 * time.Second
)

func GetTaskStorage(logger slog.Logger) (TaskStorage, error) {
	fn := "internal.db.db.GetStorage"
	db_host := os.Getenv("DB_HOST")
	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	postgres_db := os.Getenv("POSTGRES_DB")
	db_port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", postgres_user, postgres_password, db_host, db_port, postgres_db)

	var db *sql.DB
	var err error

	for i := range DATABASE_CONNECTION_TRIES_COUNT {
		// Connection with data base
		db, err = sql.Open("pgx", dsn)

		if err != nil {
			logger.Warn(fmt.Sprintf("Cannot open database(retrying in 1 second), try number: %d/%d", i+1, DATABASE_CONNECTION_TRIES_COUNT), "source", fn)
			time.Sleep(RETRY_TIME_BETWEEN_TRIES)
			continue
		}

		if err := db.Ping(); err != nil {
			logger.Warn(fmt.Sprintf("Cannot ping database, try number: %d/%d", i+1, DATABASE_CONNECTION_TRIES_COUNT), "source", fn)
			time.Sleep(RETRY_TIME_BETWEEN_TRIES)
			continue
		}
		logger.Info("esablished connection with database")
		break
	}

	if err != nil {
		return TaskStorage{}, fmt.Errorf("%s:%w", fn, err)
	}

	storage := TaskStorage{
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
	CREATE INDEX IF NOT EXISTS idx_tasks_id ON tasks(id);
	`)

	if err != nil {
		return TaskStorage{}, fmt.Errorf("cannot create notes table: %s", err.Error())
	}

	return storage, nil
}
