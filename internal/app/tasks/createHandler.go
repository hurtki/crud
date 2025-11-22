package tasksHandler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/hurtki/crud/internal/domain/tasks"
)

func (h *TasksHandler) HandleCreate(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.tasks.TasksHandler.HandleCreate"

	reqBody, err := io.ReadAll(req.Body)

	if err != nil {
		h.logger.Error("cannot read request body", "source", fn, "err", err)
		writeJSONError(res, http.StatusBadRequest, "cannot read request body")
		return
	}

	reqDto := CreateTaskRequest{}

	if err := json.Unmarshal(reqBody, &reqDto); err != nil {
		writeJSONError(res, http.StatusBadRequest, "cannot deserialize request")
		return
	}

	task, err := h.useCases.CreateTask(reqDto.ToUseCaseTask())

	if err != nil {
		var errFieldCannotBeEmpty *tasks.ErrFieldCannotBeEmpty
		if errors.As(err, &errFieldCannotBeEmpty) {
			writeJSONError(res, http.StatusBadRequest, err.Error())
			return
		}
		var errTaskAlreadyExists *tasks.ErrTaskWithThisValueAlreadyExists
		if errors.As(err, &errTaskAlreadyExists) {
			writeJSONError(res, http.StatusConflict, err.Error())
			return
		}
		if err == tasks.ErrCannotCreateTask {
			writeJSONError(res, http.StatusInternalServerError, err.Error())
			return
		}
		h.logger.Error("not handled error from usecase", "source", fn, "err", err)
		writeJSONError(res, http.StatusInternalServerError, "cannot create task")
		return
	}

	resDto := TaskResponse{
		Id:   task.Id,
		Name: task.Name,
		Text: task.Text,
	}

	resBody, err := json.Marshal(resDto)

	if err != nil {
		h.logger.Error("Cannot serialize response", "source", fn, "err", err)
		writeJSONError(res, http.StatusInternalServerError, "cannot serialize response")
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write(resBody)
}
