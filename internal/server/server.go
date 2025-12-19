package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/hurtki/crud/internal/config"
)

type Server struct {
	httpServer *http.Server
	appConfig  config.AppConfig
}

type Middleware func(http.Handler) http.Handler

// NewServer creates new server, router will be called after all middlewares
// first given middleware will be called first / last will be called last
func NewServer(router http.Handler, config config.AppConfig, middlewares ...Middleware) *Server {

	httpServer := http.Server{Addr: config.InternalPort}

	// going through all the given middlewares from the end
	// so the first given will be called first
	h := router
	for i := len(middlewares) - 1; i >= 0; i-- {
		// wrapping previous with new
		h = middlewares[i](h)
	}

	httpServer.Handler = h

	return &Server{
		appConfig:  config,
		httpServer: &httpServer,
	}
}

// Start starts the server
// errChan parameter is used to handle errors that can occur in http.Server.ListenAndServe
// ErrServerClosed is not returned to this chanel
// the best way to use with select:
/*
select {
case err = <-srvErrChan:
	logger.Error("error occured in server", "err", err)
case <-signalChan:
	logger.Info("Stopping server...")
	err = srv.Stop()
}
*/
func (s *Server) Start(errChan chan error) {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()
}

// Stop unexpectedly stops server using timeout context
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.appConfig.ServerTimeToShutDown)

	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
