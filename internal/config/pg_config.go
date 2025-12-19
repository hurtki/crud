package config

import "os"

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

func LoadPgConfig() *PostgresConfig {
	pgConfig := PostgresConfig{}

	pgConfig.Host = os.Getenv("DB_HOST")
	pgConfig.User = os.Getenv("POSTGRES_USER")
	pgConfig.Password = os.Getenv("POSTGRES_PASSWORD")
	pgConfig.DbName = os.Getenv("POSTGRES_DB")
	pgConfig.Port = os.Getenv("DB_PORT")

	return &pgConfig
}
