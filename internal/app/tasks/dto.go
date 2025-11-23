package tasksHandler

type CreateTaskRequest struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type PutTaskRequest struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type TaskResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

type ListTaskResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}
