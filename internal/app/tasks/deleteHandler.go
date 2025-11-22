package tasksHandler

import (
	"net/http"

	"github.com/hurtki/crud/internal/domain/tasks"
)

func (h *TasksHandler) HandleDelete(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.tasks.TasksHandler.HandleDelete"

	idUrlParameter, ok := req.Context().Value("urlParameter").(int)
	if !ok {
		h.logger.Error("cannnot get int id url parameter from context", "source", fn)
		http.Error(res, "cannot get url parameter with id", http.StatusInternalServerError)
		return
	}

	err := h.useCases.DeleteTask(ToUseCaseTaskDelete(idUrlParameter))

	if err != nil {
		if err == tasks.ErrTaskIdSmallerThanNull {
			writeJSONError(res, http.StatusBadRequest, "id should be bigger than null")
			return
		}
		if err == tasks.ErrTaskWithIdNotFound {
			writeJSONError(res, http.StatusNotFound, "task with given id not found")
			return
		}
		if err == tasks.ErrCannotDeleteTask {
			writeJSONError(res, http.StatusInternalServerError, err.Error())
			return
		}
		h.logger.Error("not handled error from usecase", "source", fn, "err", err)
		writeJSONError(res, http.StatusInternalServerError, "cannot delete task")
		return
	}

	res.WriteHeader(http.StatusNoContent)
}
