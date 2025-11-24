package tasksHandler

import (
	"encoding/json"
	"net/http"

	domain_tasks "github.com/hurtki/crud/internal/domain/tasks"
)

func (h *TasksHandler) HandleList(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.tasks.getHandler.HandleList"

	tasks, err := h.useCases.GetTasks()
	if err != nil {
		if err == domain_tasks.ErrCannotGetTasks {
			writeJSONError(res, http.StatusInternalServerError, err.Error())
			return
		}
		h.logger.Error("not handled error from usecase", "source", fn, "err", err)
		writeJSONError(res, http.StatusInternalServerError, "cannot list tasks")
		return
	}
	serialized, err := json.Marshal(tasks)

	if err != nil {
		h.logger.Error("error serializing tasks in to json", "source", fn, "err", err)
		writeJSONError(res, http.StatusInternalServerError, "cannot serialize tasks to json")
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(serialized)
}
