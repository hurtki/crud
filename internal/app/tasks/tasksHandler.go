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

func (h *TasksHandler) ServeReadUpdateDelete(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.HandleRead(res, req)
	case "PUT", "PATCH":
		h.HandleUpdate(res, req)
	case "DELETE":
		h.HandleDelete(res, req)
	default:
		http.NotFound(res, req)
	}
}

func (h *TasksHandler) ServeCreateList(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.HandleList(res, req)
	case "POST":
		h.HandleCreate(res, req)
	default:
		http.NotFound(res, req)
	}
}
