package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	internalAppPort = ":80"
)

type AppConfig struct {
	InternalPort         string        // :port_id
	TasksPerPageCount    int           `yaml:"tasks_per_page"`    // count of tasks that will be on list endpoint on every page ( limit )
	ServerTimeToShutDown time.Duration `yaml:"time_to_shut_down"` // time to wait to close http server, when gracefull shutdown intialized
	LoggingLevel         slog.Level    `yaml:"logging_level"`     // logging level, minimal level that will be logged
	Cors                 bool          `yaml:"use_cors"`          // if the server will use cors midlleware
	CorsOrigins          []string      `yaml:"cors_origins"`      // if cors is on, what origins will it contain
}

func LoadConfig(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.TasksPerPageCount < 1 {
		return nil, ErrTasksPerPageSmallerThanOne
	}

	if cfg.ServerTimeToShutDown <= 0 {
		return nil, ErrShutdownTimeSmallerThanNull
	}

	if cfg.Cors && len(cfg.CorsOrigins) == 0 {
		return nil, ErrNoOriginsWithCors
	}

	cfg.InternalPort = internalAppPort

	return &cfg, nil
}

func (c AppConfig) String() string {
	return fmt.Sprintf(
		"AppConfig: "+
			" LoggingLevel: %s"+
			"  InternalPort ( in container ): %s "+
			"  TasksPerPageCount: %d "+
			"  ServerTimeToShutDown: %s "+
			"  Cors: %t "+
			"  CorsOrigins: %v ",
		c.LoggingLevel,
		c.InternalPort,
		c.TasksPerPageCount,
		c.ServerTimeToShutDown,
		c.Cors,
		c.CorsOrigins,
	)
}
