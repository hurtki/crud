package tasksHandler

import (
	"net/http"
	"encoding/json"
)

func (h *TasksHandler) HandleGet(res http.ResponseWriter, req *http.Request) {
	h.logger.Info("IN GET")
	tasks, err := h.storage.GetTasks()
	if err != nil {
		res.WriteHeader(500)
		return
	}
	data, _ := json.Marshal(tasks)	

	res.Write([]byte(data))
	res.WriteHeader(200)
}
