package tasksHandler

import (
	"fmt"
	"net/http"
)

func (h *TasksHandler) HandleUpdate(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.tasks.TasksHandler.HandleUpdate"
	h.logger.Info(fmt.Sprintf("got parameter from upper: %v", req.Context().Value("urlParameter")))

	h.logger.Warn("not filled function", "source", fn)
	res.WriteHeader(300)
}
