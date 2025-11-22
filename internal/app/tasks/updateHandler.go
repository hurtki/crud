package tasksHandler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/hurtki/crud/internal/domain/tasks"
)

func (h *TasksHandler) HandleUpdate(res http.ResponseWriter, req *http.Request) {
	fn := "internal.app.tasks.TasksHandler.HandleUpdate"

	idUrlParameter, ok := req.Context().Value("urlParameter").(int)
	if !ok {
		h.logger.Error("cannnot get int id url parameter from context", "source", fn)
		http.Error(res, "cannot get url parameter with id", http.StatusInternalServerError)
		return
	}

	reqBody, err := io.ReadAll(req.Body)

	if err != nil {
		h.logger.Error("cannot read request body", "source", fn, "err", err)
		writeJSONError(res, http.StatusBadRequest, "cannot read request body")
		return
	}

	reqDto := PutTaskRequest{}

	if err := json.Unmarshal(reqBody, &reqDto); err != nil {
		writeJSONError(res, http.StatusBadRequest, "cannot deserialize request")
		return
	}

	useCaseInput := tasks.UpdateTaskInput{
		Id:   idUrlParameter,
		Name: reqDto.Name,
		Text: reqDto.Text,
	}

	task, err := h.useCases.UpdateTask(useCaseInput)

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
		if err == tasks.ErrTaskWithIdNotFound {
			writeJSONError(res, http.StatusBadRequest, err.Error())
			return
		}
		if err == tasks.ErrTaskIdSmallerThanNull {
			writeJSONError(res, http.StatusBadRequest, err.Error())
			return
		}
		if err == tasks.ErrCannotUpdateTask {
			writeJSONError(res, http.StatusInternalServerError, err.Error())
			return
		}

		h.logger.Error("not handled error from usecase", "source", fn, "err", err)
		writeJSONError(res, http.StatusInternalServerError, "cannot update task")
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
