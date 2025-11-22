package tasks

type TaskRepository interface {
	Create(task Task) (Task, error)
	Update(task Task) (Task, error)
	GetByID(id int) (Task, error)
	List() ([]Task, error)
	Delete(id int) error
}
