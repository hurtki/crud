package tasksHandler


import (
	"net/http"
	"fmt"
)

func (h *TasksHandler) HandleRead(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.tasks.TasksHandler.HandleRead"
	h.logger.Info(fmt.Sprintf("got parameter from upper: %v", req.Context().Value("urlParameter")))

	h.logger.Warn("not filled function", "source", fn)
	res.WriteHeader(300)
}


