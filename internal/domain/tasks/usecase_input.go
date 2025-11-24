package tasks

type CreateTaskInput struct {
	Name string
	Text string
}

type GetTaskInput struct {
	Id int
}

type UpdateTaskInput struct {
	Id   int
	Name string
	Text string
}

type DeleteTaskInput struct {
	Id int
}
