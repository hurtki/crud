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
		logger: logger.With("service", "root_handler"),
		storage: storage,
	}
}

func (h *TasksHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.logger.Info("started serving in tasks handler")
	switch req.Method {
	case "GET":
		h.HandleGet(res, req)
	case "POST":
		h.HandlePost(res, req)
	}
}
