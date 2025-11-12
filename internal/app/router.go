package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/hurtki/crud/internal/config"
)

type Router struct {
	logger   *slog.Logger
	config   config.AppConfig
	routeSet RouteSet
}

func NewRouter(logger slog.Logger, cgf config.AppConfig, routeSet RouteSet) Router {

	router := Router{
		logger:   logger.With("service", "HTTP-Hanlder"),
		config:   cgf,
		routeSet: routeSet,
	}

	return router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.logger.Info("handling request", "path", req.URL.Path, "method", req.Method, "ip", req.RemoteAddr)

	handler, err := r.routeSet.Handler(req.URL.Path)
	if err != nil {
		http.NotFound(w, req)
		return
	}

	handler.ServeHTTP(w, req)
}

func (r *Router) StartRouting() error {
	r.logger.Info(fmt.Sprintf("started routing on port: %s", r.config.Port))
	return http.ListenAndServe(r.config.Port, r)
}
