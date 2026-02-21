package tasksHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	domain_tasks "github.com/hurtki/crud/internal/domain/tasks"
)

func (h *TasksHandler) HandleList(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.tasks.getHandler.HandleList"

	var pageNumber int
	var err error
	pageParameter := req.URL.Query().Get("page")
	if pageParameter == "" {
		pageNumber = 0
	} else {
		pageNumber, err = strconv.Atoi(pageParameter)
		if err != nil {
			writeJSONError(res, http.StatusBadRequest, "page parameter should be int")
			return
		}
	}

	tasks, err := h.useCases.ListTasks(pageNumber)
	if err != nil {
		if err == domain_tasks.ErrCannotGetTasks {
			writeJSONError(res, http.StatusInternalServerError, err.Error())
			return
		}
		if err == domain_tasks.ErrPageCannotBeSmallerThanNull {
			writeJSONError(res, http.StatusBadRequest, err.Error())
			return
		}
		h.logger.Error("not handled error from usecase", "source", fn, "err", err)
		writeJSONError(res, http.StatusInternalServerError, "cannot list tasks")
		return
	}
	resDto := make([]TaskResponse, len(tasks))
	for i := range len(tasks) {
		resDto[i] = TaskResponse{
			Id:   tasks[i].Id,
			Name: tasks[i].Name,
			Text: tasks[i].Text,
		}

	}
	serialized, err := json.Marshal(resDto)

	if err != nil {
		h.logger.Error("error serializing tasks in to json", "source", fn, "err", err)
		writeJSONError(res, http.StatusInternalServerError, "cannot serialize tasks to json")
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(serialized)
}
