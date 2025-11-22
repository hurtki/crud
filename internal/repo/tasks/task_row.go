package tasks_repo

import (
	"github.com/hurtki/crud/internal/domain/tasks"
)

// TODO: use only this structure from in db operations
type TaskRow struct {
	Id   int
	Name string
	Text string
}

func ToDomainTask(tr TaskRow) tasks.Task {
	return tasks.NewTask(
		tr.Id,
		tr.Name,
		tr.Text,
	)
}
