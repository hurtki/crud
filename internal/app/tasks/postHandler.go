package tasksHandler

import (
	"net/http"
)

func (h *TasksHandler) HandlePost(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(300)
}
