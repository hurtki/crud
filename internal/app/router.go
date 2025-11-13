package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/hurtki/crud/internal/config"
)

// Router is a structure to handle requests and routing them to handlers
// Router uses RouteSet to match path with handler
type Router struct {
	logger   *slog.Logger
	config   config.AppConfig
	routeSet RouteSet
}

func NewRouter(logger slog.Logger, cgf config.AppConfig, routeSet RouteSet) Router {

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

	handler, err := r.routeSet.Handler(req.URL.Path)
	if err == ErrNotFound {
		http.NotFound(res, req)
		return
	} else if err != nil {
		r.logger.Error("unexpected error, when getting handler for path", "source", fn)
		// even with error also sending, that we cannot found the route
		http.NotFound(res, req)
	}
	
	// Serving with, appropriate to route, handler 
	handler.ServeHTTP(res, req)
}

func (r *Router) StartRouting() error {
	r.logger.Info(fmt.Sprintf("started routing on port: %s", r.config.Port))
	return http.ListenAndServe(r.config.Port, r)
}
