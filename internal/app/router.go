package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/hurtki/crud/internal/app/routeSet"

	"github.com/hurtki/crud/internal/config"
)

// Router is a structure to handle requests and routing them to handlers
// Router uses RouteSet to match path with handler
type Router struct {
	logger   *slog.Logger
	config   config.AppConfig
	routeSet routeSet.RouteSet
}

func NewRouter(logger slog.Logger, cgf config.AppConfig, routeSet routeSet.RouteSet) Router {

	router := Router{
		// wrap of logger with "service" field
		logger:   logger.With("service", "HTTP-Hanlder"),
		config:   cgf,
		routeSet: routeSet,
	}

	return router
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.router.Router.ServeHTTP"

	r.logger.Info("INCOMING REQUEST:", "path", req.URL.Path, "method", req.Method, "ip", req.RemoteAddr)

	handler, parameter, err := r.routeSet.Handler(req.URL.Path)
	if err == routeSet.ErrNotFound {
		http.NotFound(res, req)
		return
	} else if err != nil {
		r.logger.Error("unexpected error, when getting handler for path", "source", fn)
		// even with error also sending, that we cannot found the route
		http.NotFound(res, req)
	}

	if parameter != nil {
		ctx := context.WithValue(req.Context(), "urlParameter", parameter)
		req = req.WithContext(ctx)
	}

	// Serving with, appropriate to route, handler
	handler.ServeHTTP(res, req)
}

// StartRouting starts listening in AppConfig's port and routing them to specified in routeSet handlers
func (r *Router) StartRouting() error {
	r.logger.Info(fmt.Sprintf("started routing on port: %s", r.config.Port))
	return http.ListenAndServe(r.config.Port, r)
}
