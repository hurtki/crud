package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/hurtki/crud/internal/server"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware first runs all next Handlers and then logs result information about done request ( with satatus code )
func LoggingMiddleware(logger *slog.Logger) server.Middleware {
	logger = logger.With("service", "response-logging-middleware")

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := &responseWriter{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			before := time.Now()
			next.ServeHTTP(rw, r)

			logger.Info("request completed",
				"method", r.Method,
				"path", r.URL.Path,
				"status", rw.status,
				"duration_ms", time.Since(before).Milliseconds(),
			)
		})

	}

}
