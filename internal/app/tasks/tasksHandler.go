package tasksHandler

import (
	"log/slog"
	"net/http"

	"github.com/hurtki/crud/internal/domain/tasks"
)

type TasksHandler struct {
	logger   *slog.Logger
	useCases tasks.TaskUseCases
}

func NewTasksHandler(logger slog.Logger, useCases tasks.TaskUseCases) TasksHandler {
	return TasksHandler{
		// wrap of logger with "serivice" field
		logger:   logger.With("service", "tasksHandler"),
		useCases: useCases,
	}
}

func (h *TasksHandler) ServeReadUpdateDelete(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.HandleRead(res, req)
	case "PUT":
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
