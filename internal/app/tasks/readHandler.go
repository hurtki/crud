package tasksHandler

import (
	"encoding/json"
	"net/http"

	"github.com/hurtki/crud/internal/domain/tasks"
)

func (h *TasksHandler) HandleRead(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.tasks.TasksHandler.HandleRead"

	idUrlParameter, ok := req.Context().Value("urlParameter").(int)
	if !ok {
		h.logger.Error("cannnot get int id url parameter from context", "source", fn)
		http.Error(res, "cannot get url parameter with id", http.StatusInternalServerError)
		return
	}

	useCaseInput := tasks.GetTaskInput{Id: idUrlParameter}
	task, err := h.useCases.GetTask(useCaseInput)
	if err != nil {
		if err == tasks.ErrTaskIdSmallerThanNull {
			writeJSONError(res, http.StatusBadRequest, "id should be bigger than null")
			return
		}
		if err == tasks.ErrTaskWithIdNotFound {
			writeJSONError(res, http.StatusNotFound, "task with given id not found")
			return
		}
		if err == tasks.ErrCannotGetTask {
			writeJSONError(res, http.StatusInternalServerError, err.Error())
			return
		}
		h.logger.Error("not handled error from usecase", "source", fn, "err", err)
		writeJSONError(res, http.StatusInternalServerError, "cannot delete task")
		return
	}

	serialized, err := json.Marshal(task)

	if err != nil {
		h.logger.Error("error serializing task into json", "source", fn, "err", err)
		writeJSONError(res, http.StatusInternalServerError, "cannot serialize task to json")
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(serialized)
}
