package tasksHandler

import (
	"net/http"
)

func (h *TasksHandler) HandleCreate(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.tasks.TasksHandler.HandleCreate"

	h.logger.Warn("not filled function", "source", fn)
	res.WriteHeader(300)
}
