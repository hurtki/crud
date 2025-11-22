package tasks

import (
	"errors"

	repoerr "github.com/hurtki/crud/internal/repo"
)

type TaskUseCases struct {
	repo TaskRepository
}

func NewTaskUseCases(repo TaskRepository) TaskUseCases {
	return TaskUseCases{repo: repo}
}

func (c *TaskUseCases) CreateTask(input CreateTaskInput) (Task, error) {
	if input.Name == "" {
		return Task{}, &ErrFieldCannotBeEmpty{Field: "name"}
	}
	task := NewTask(0, input.Name, input.Text)

	task, err := c.repo.Create(task)
	if err != nil {
		var errEmptyField *repoerr.ErrEmptyField
		var errRepoInternal *repoerr.ErrRepoInternal

		switch {
		case errors.As(err, &errEmptyField):
			err = &ErrFieldCannotBeEmpty{Field: errEmptyField.Field}
		case errors.As(err, &errRepoInternal):
			err = ErrCannotCreateTask
		default:
			err = ErrCannotCreateTask
		}
		return Task{}, err
	}

	return task, nil
}

func (c *TaskUseCases) GetTask(input GetTaskInput) (Task, error) {
	if input.Id < 1 {
		return Task{}, ErrTaskIdSmallerThanNull
	}
	task, err := c.repo.GetByID(input.Id)

	if err != nil {

		if err == repoerr.ErrNothingFound {

			return Task{}, ErrTaskWithIdNotFound
		}
		var errRepoInternal *repoerr.ErrRepoInternal
		if errors.As(err, &errRepoInternal) {
			return Task{}, ErrCannotGetTask
		}
		return Task{}, ErrCannotGetTask
	}

	return task, nil
}

func (c *TaskUseCases) GetTasks() ([]Task, error) {
	tasks, err := c.repo.List()

	if err != nil {
		var errRepoInternal *repoerr.ErrRepoInternal
		if errors.As(err, &errRepoInternal) {
			return []Task{}, ErrCannotGetTasks
		}
		return []Task{}, ErrCannotGetTasks
	}

	return tasks, nil
}

func (c *TaskUseCases) UpdateTask(input UpdateTaskInput) (Task, error) {
	if input.Id < 1 {
		return Task{}, ErrTaskIdSmallerThanNull
	}
	if input.Name == "" {
		return Task{}, &ErrFieldCannotBeEmpty{Field: "name"}
	}
	task := NewTask(input.Id, input.Name, input.Text)
	task, err := c.repo.Update(task)
	if err != nil {
		var errConflictValue *repoerr.ErrConflictValue
		var errEmptyField *repoerr.ErrEmptyField
		var errRepoInternal *repoerr.ErrRepoInternal

		switch {
		case errors.As(err, &errConflictValue):
			err = &ErrTaskWithThisValueAlreadyExists{Field: errConflictValue.Field}
		case errors.As(err, &errEmptyField):
			err = &ErrFieldCannotBeEmpty{Field: errEmptyField.Field}
		case errors.As(err, &errRepoInternal):
			err = ErrCannotUpdateTask
		// TODO, need to stay only one of this to repo errors, very strange logic
		case err == repoerr.ErrNothingChanged || err == repoerr.ErrNothingFound:
			err = ErrTaskWithIdNotFound
		default:
			err = ErrCannotUpdateTask
		}
		return Task{}, err
	}

	return task, nil
}

func (c *TaskUseCases) DeleteTask(input DeleteTaskInput) error {
	if input.Id < 1 {
		return ErrTaskIdSmallerThanNull
	}
	err := c.repo.Delete(input.Id)

	if err != nil {
		// TODO, need to stay only one of this to repo errors, very strange logic
		if err == repoerr.ErrNothingFound || err == repoerr.ErrNothingChanged {
			return ErrTaskWithIdNotFound
		}

		var errRepoInternal *repoerr.ErrRepoInternal
		if errors.As(err, &errRepoInternal) {
			return ErrCannotDeleteTask
		}
		return ErrCannotDeleteTask
	}
	return nil
}
