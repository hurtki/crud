package tasks

type TaskRepository interface {
	Create(task Task) (Task, error)
	Update(task Task) (Task, error)
	GetByID(id int) (Task, error)
	List(Pagination) ([]Task, error)
	Delete(id int) error
}

type Pagination struct {
	Limit  int
	Cursor int
}
