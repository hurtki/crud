package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/hurtki/crud/internal/config"
	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
	appConfig  config.AppConfig
}

func NewServer(router http.Handler, config config.AppConfig) *Server {
	httpServer := http.Server{Addr: ":80"}

	if config.Cors {
		allowedMethods := []string{"GET", "POST", "PUT", "DELETE"}
		c := cors.New(cors.Options{AllowedOrigins: config.CorsOrigins, AllowCredentials: true, AllowedMethods: allowedMethods})
		httpServer.Handler = c.Handler(router)
	} else {
		httpServer.Handler = router
	}

	return &Server{
		appConfig:  config,
		httpServer: &httpServer,
	}
}

func (s *Server) Start(errChan chan error) {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.appConfig.ServerTimeToShutDown)

	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
