package config

import (
	"regexp"
)

// NewAppConfig creates a new AppConfig entity with validation
// all AppConfig fields and how they should look see in AppConfig structure
func NewAppConfig(port string) AppConfig {
	port_reg_exp := regexp.MustCompile(`^:\d{1,5}$`)
	if !port_reg_exp.MatchString(port) {
		panic("wrong port specified")
	}

	return AppConfig{
		Port: port,
	}
}

type AppConfig struct {
	Port string // :port_id
}
