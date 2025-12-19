package middleware

import (
	"net/http"

	"github.com/hurtki/crud/internal/config"
	"github.com/hurtki/crud/internal/server"
	"github.com/rs/cors"
)

// CorstMiddleware adds cors headers according to given appConfig and specified in it origins
func CorsMiddleware(appConfig *config.AppConfig) server.Middleware {
	var h http.Handler
	return func(next http.Handler) http.Handler {

		if appConfig.Cors {
			allowedMethods := []string{"GET", "POST", "PUT", "DELETE"}
			c := cors.New(cors.Options{AllowedOrigins: appConfig.CorsOrigins, AllowCredentials: true, AllowedMethods: allowedMethods})
			h = c.Handler(next)
		} else {
			h = next
		}

		return h

	}
}
