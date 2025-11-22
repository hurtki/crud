package tasksHandler

import (
	"github.com/hurtki/crud/internal/domain/tasks"
)

func ToTasksReponse(tasks []tasks.Task) ListTaskResponse {
	lstTaskResponse := ListTaskResponse{Tasks: []TaskResponse{}}
	for _, task := range tasks {
		lstTaskResponse.Tasks = append(lstTaskResponse.Tasks, TaskResponse{
			Id:   task.Id,
			Name: task.Name,
			Text: task.Text,
		})
	}
	return lstTaskResponse
}

func (r *CreateTaskRequest) ToUseCaseTask() tasks.CreateTaskInput {
	return tasks.CreateTaskInput{
		Name: r.Name,
		Text: r.Text,
	}
}

func ToUseCaseTaskDelete(id int) tasks.DeleteTaskInput {
	return tasks.DeleteTaskInput{
		Id: id,
	}
}
