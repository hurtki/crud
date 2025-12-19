package tasks_repo

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/hurtki/crud/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

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

func GetDb(logger *slog.Logger, pgConf *config.PostgresConfig) (*sql.DB, error) {
	fn := "internal.repo.tasks.GetDb"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pgConf.User, pgConf.Password, pgConf.Host, pgConf.Port, pgConf.DbName)

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
			logger.Warn(fmt.Sprintf("Cannot ping database, try number: %d/%d", i+1, dataBaseConnectionTriesCount), "source", fn, "err", err)
			time.Sleep(retryTimeBetweenTries)
			continue
		}
		logger.Info("esablished connection with database")
		break
	}

	if err != nil {
		logger.Error("Can't establish connection with database", "source", fn, "err", err)
		return nil, fmt.Errorf("%s:%w", fn, err)
	}

	maxConnectionsRow := db.QueryRow(`
	SHOW max_connections;
	`)
	var maxPostgresConnections int
	if err := maxConnectionsRow.Scan(&maxPostgresConnections); err != nil {
		logger.Error("cannot get postgres max connections", "err", err, "source", fn)
		return nil, fmt.Errorf("can't get postgres max connections, can't intialize storage: %w", err)
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
	return db, nil
}
