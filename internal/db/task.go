package db

import (
	"fmt"

	"github.com/hurtki/crud/internal/domain/tasks"
)

func (s *Storage) GetTasks() ([]tasks.Task, error) {
	fn := "internal.db.task.GetTasks"
	res, err := s.db.Query(`
	SELECT * FROM tasks;
	`)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", fn, err)
	}
	task := tasks.Task{}
	tasks := []tasks.Task{}

	for res.Next() {

		if err := res.Scan(&task.Id, &task.Name, &task.Text); err != nil {
			return nil, fmt.Errorf("%s:%w", fn, err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Storage) AddTask(task tasks.Task) error {
	fn := "internal.db.task.AddTask"

	res, err := s.db.Exec(`
	INSERT INTO tasks (name, text)
	VALUES ($1, $2);
	`, task.Name, task.Text)

	if err != nil {
		return fmt.Errorf("%s:%w", fn, err)
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected == 0 {
		return fmt.Errorf("%s:%w", fn, ErrorNoRowsAffected)
	}

	return nil
}
