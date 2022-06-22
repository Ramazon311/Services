package models

type Task struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Assignee_id string `json:"Assignee_id"`
	Status      string `json:"status"`
	Deadline    string `json:"deadline"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
	Deleted_at  string `json:"deleted_at"`
}


type CreateTask struct {
	Assignee_id string `json:"assignee_id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Status      string `json:"status"`
	Deadline    string `json:"deadline"`
}


type UpdateTasks struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Assignee_id string `json:"assignee_id"`
	Status      string `json:"status"`
	Deadline    string `json:"deadline"`
}
