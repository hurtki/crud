package tasks_repo

import (
	"errors"

	"github.com/hurtki/crud/internal/domain/tasks"
	repoerr "github.com/hurtki/crud/internal/repo"
)

func (s *TaskStorage) List(pag tasks.Pagination) ([]tasks.Task, error) {
	fn := "internal.repo.tasks.TaskStorage.List"
	res, err := s.db.Query(`
	SELECT id, name, text FROM tasks
	ORDER BY id ASC
	LIMIT $1 OFFSET $2;
	`, pag.Limit, pag.Cursor)

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

func (s *TaskStorage) Create(task tasks.Task) (tasks.Task, error) {
	fn := "internal.repo.tasks.TaskStorage.Create"

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

func (s *TaskStorage) Update(task tasks.Task) (tasks.Task, error) {
	fn := "internal.repo.tasks.TaskStorage.Update"

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
		return tasks.Task{}, repoerr.ErrNothingChanged
	}

	return task, nil
}

func (s *TaskStorage) GetByID(id int) (tasks.Task, error) {
	fn := "internal.repo.tasks.TaskStorage.GetByID"

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

func (s *TaskStorage) Delete(id int) error {
	fn := "internal.repo.tasks.TaskStorage.Delete"

	res, err := s.db.Exec(`
	DELETE FROM tasks
	WHERE id = $1
	`, id)
	if err != nil {
		return s.handleDbErr(fn, err)
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected < 1 {
		return repoerr.ErrNothingChanged
	}

	return nil
}

func (s *TaskStorage) handleDbErr(fn string, err error) error {
	repoErr := toRepoError(err)
	var internalErr *repoerr.ErrRepoInternal
	if errors.As(repoErr, &internalErr) {
		s.logger.Error("repo internal error", "source", fn, "err", repoErr)
	}
	return repoErr
}
