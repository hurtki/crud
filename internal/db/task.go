package db

import (
	"errors"

	"github.com/hurtki/crud/internal/domain/tasks"
)

func (s *Storage) List() ([]tasks.Task, error) {
	fn := "internal.db.task.Storage.List"
	res, err := s.db.Query(`
	SELECT id, name, text FROM tasks;
	`)

	if err != nil {
		return nil, s.handleDbErr(fn, err)
	}
	tasksSlice := []tasks.Task{}

	for res.Next() {
		task := tasks.Task{}
		if err := res.Scan(&task.Id, &task.Name, &task.Text); err != nil {
			s.logger.Error("unexcpected error from scan", "source", fn, "err", err)
			return nil, err
		}
		tasksSlice = append(tasksSlice, task)
	}

	if err := res.Err(); err != nil {
		s.logger.Error("error after iterating by rows", "source", fn, "err", err)
		return nil, err
	}

	return tasksSlice, nil
}

func (s *Storage) Create(task tasks.Task) (tasks.Task, error) {
	fn := "internal.db.task.Storage.Create"

	row := s.db.QueryRow(`
	INSERT INTO tasks (name, text)
	VALUES ($1, $2)
	RETURNING id;
	`, task.Name, task.Text)

	err := row.Scan(&task.Id)

	if err != nil {
		return tasks.Task{}, s.handleDbErr(fn, err)
	}

	return task, nil
}

func (s *Storage) Update(task tasks.Task) (tasks.Task, error) {
	fn := "internal.db.task.Storage.Update"

	res, err := s.db.Exec(`
	UPDATE tasks
	SET name = $2,
		text = $3
	WHERE id = $1;
	`, task.Id, task.Name, task.Text,
	)

	if err != nil {
		return tasks.Task{}, s.handleDbErr(fn, err)
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected < 1 {
		return tasks.Task{}, ErrNoRowsAffected
	}

	return task, nil
}

func (s *Storage) GetByID(id int) (tasks.Task, error) {
	fn := "internal.db.task.Storage.GetByID"

	row := s.db.QueryRow(`
	SELECT id, name, text FROM tasks
	WHERE id = $1;
	`, id)

	var task tasks.Task
	err := row.Scan(&task.Id, &task.Name, &task.Text)

	if err != nil {
		return tasks.Task{}, s.handleDbErr(fn, err)
	}

	return task, nil
}

func (s *Storage) Delete(id int) error {
	fn := "internal.db.task.Storage.Delete"
	res, err := s.db.Exec(`
	DELETE FROM tasks
	WHERE id = $1
	`, id)
	if err != nil {
		return s.handleDbErr(fn, err)
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected < 1 {
		return ErrNoRowsAffected
	}

	return nil
}

func (s *Storage) handleDbErr(fn string, err error) error {
	if dbErr, ok := ToDbError(err); ok {
		var syntaxErr *ErrSyntaxSql
		if errors.As(dbErr, &syntaxErr) {
			s.logger.Error("sql syntax error", "source", fn, "err", dbErr)
		}
		return dbErr
	}
	return err
}
