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
	// DB_MAX_CONNECTIONS_PERCENT_TO_POSTGRES_MAX - constant is a percent of db connections
	// from postgres max connections
	dbMaxConnectionsPercentToPostgresMax = 70
	// DB_MAX_IDLE_CONNECTIONS_PERCENT_TO_POSTGRES_MAX - constant is a percent of db idling connections
	// from postgres max connections
	dbMaxIdleConnectionsPercentToPostgresMax = 10

	connectionLifiTime           = 5 * time.Minute
	dataBaseConnectionTriesCount = 10
	retryTimeBetweenTries        = 1 * time.Second
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

	for i := range dataBaseConnectionTriesCount {
		// Connection with data base
		db, err = sql.Open("pgx", dsn)

		if err != nil {
			logger.Warn(fmt.Sprintf("Cannot open database(retrying in 1 second), try number: %d/%d", i+1, dataBaseConnectionTriesCount), "source", fn)
			time.Sleep(retryTimeBetweenTries)
			continue
		}

		if err := db.Ping(); err != nil {
			logger.Warn(fmt.Sprintf("Cannot ping database, try number: %d/%d", i+1, dataBaseConnectionTriesCount), "source", fn)
			time.Sleep(retryTimeBetweenTries)
			continue
		}
		logger.Info("esablished connection with database")
		break
	}

	if err != nil {
		logger.Error("Can't establish connection with database", "source", fn, "err", err)
		return TaskStorage{}, fmt.Errorf("%s:%w", fn, err)
	}

	maxConnectionsRow := db.QueryRow(`
	SHOW max_connections;
	`)
	var maxPostgresConnections int
	if err := maxConnectionsRow.Scan(&maxPostgresConnections); err != nil {
		logger.Error("cannot get postgres max connections", "err", err, "source", fn)
		return TaskStorage{}, fmt.Errorf("can't get postgres max connections, can't intialize storage: %w", err)
	}

	logger.Info("got postgres max connections count", "count", maxPostgresConnections, "source", fn)

	var dbMaxOpenCons int = maxPostgresConnections * dbMaxConnectionsPercentToPostgresMax / 100
	var dbMaxIdleCons int = maxPostgresConnections * dbMaxIdleConnectionsPercentToPostgresMax / 100

	logger.Info("setting db max open connections", "count", dbMaxOpenCons, "source", fn)
	logger.Info("setting db max idle connections", "count", dbMaxIdleCons, "source", fn)

	db.SetMaxOpenConns(dbMaxOpenCons)
	db.SetMaxIdleConns(dbMaxIdleCons)
	logger.Info("setting connection life time", "count", connectionLifiTime, "source", fn)
	db.SetConnMaxLifetime(connectionLifiTime)

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
