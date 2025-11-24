package tasks

type Task struct {
	Id   int
	Name string
	Text string
}

func NewTask(Id int, Name string, Text string) Task {
	return Task{
		Id:   Id,
		Name: Name,
		Text: Text,
	}
}
