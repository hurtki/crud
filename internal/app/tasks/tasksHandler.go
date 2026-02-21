package tasksHandler

import (
	"log/slog"

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
