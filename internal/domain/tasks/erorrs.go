package tasks

import (
	"errors"
	"fmt"
)

type ErrFieldCannotBeEmpty struct{ Field string }

func (e ErrFieldCannotBeEmpty) Error() string {
	return fmt.Sprintf("field %s should be not empty", e.Field)
}

type ErrTaskWithThisValueAlreadyExists struct{ Field string }

func (e ErrTaskWithThisValueAlreadyExists) Error() string {
	return "conflict, there is already a Task with same value on field: " + e.Field
}

var (
	ErrCannotCreateTask = errors.New("Cannot create task")
	ErrCannotGetTask    = errors.New("Cannot get task")
	ErrCannotGetTasks   = errors.New("Cannot get tasks")
	ErrCannotUpdateTask = errors.New("cannot update task")
	ErrCannotDeleteTask = errors.New("cannot delete task")
	
	ErrPageCannotBeSmallerThanNull = errors.New("page cannot be smaller than null")
	ErrTaskWithIdNotFound    = errors.New("task with given id not found")
	ErrTaskIdSmallerThanNull = errors.New("Task id smaller than null")
)
