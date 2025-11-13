package tasksHandler

import (
	"log/slog"
	"net/http"

	"github.com/hurtki/crud/internal/config"
	"github.com/hurtki/crud/internal/db"
)

type TasksHandler struct {
	config  config.AppConfig
	logger  *slog.Logger
	storage *db.Storage
}

func NewTasksHandler(logger slog.Logger, config config.AppConfig, storage *db.Storage) TasksHandler {
	return TasksHandler{
		config: config,
		// wrap of logger with "serivice" field
		logger:  logger.With("service", "tasksHandler"),
		storage: storage,
	}
}

func (h *TasksHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.HandleGet(res, req)
	case "POST":
		h.HandlePost(res, req)
	}
}
