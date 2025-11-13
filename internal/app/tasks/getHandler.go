package tasksHandler

import (
	"encoding/json"
	"net/http"
)

func (h *TasksHandler) HandleGet(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.tasks.getHandler.HandleGet"
	tasks, err := h.storage.GetTasks()
	if err != nil {
		h.logger.Error("error from database, when getting tasks", "source", fn)
		http.Error(res, "error getting tasks", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(tasks)

	if err != nil {
		h.logger.Error("error while marshaling tasks to json", "source", fn)
		http.Error(res, "failder to encode tasks", http.StatusInternalServerError)
		return
	}

	res.Write([]byte(data))
}
