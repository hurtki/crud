package config

import (
	"regexp"
	"time"
)

type AppConfig struct {
	Port                 string        // :port_id
	TasksPerPageCount    int           // count of tasks that will be on list endpoint on every page ( limit )
	ServerTimeToShutDown time.Duration // time to wait to close http server, when gracefull shutdown intialized
}

// NewAppConfig creates a new AppConfig entity with validation
// all AppConfig fields and how they should look see in AppConfig structure
func NewAppConfig(port string, tasksPerPageCount int, serverTimeToShutDown time.Duration) *AppConfig {
	port_reg_exp := regexp.MustCompile(`^:\d{1,5}$`)
	if !port_reg_exp.MatchString(port) {
		panic("wrong port specified")
	}

	if tasksPerPageCount < 1 {
		panic("tasks per page count should be bigger than null")
	}

	return &AppConfig{
		Port:                 port,
		TasksPerPageCount:    tasksPerPageCount,
		ServerTimeToShutDown: serverTimeToShutDown,
	}
}
