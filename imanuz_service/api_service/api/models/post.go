package models

type Post struct {
	Id         string `json:"id"`
	User_id    string `json:"user_id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}

type UpdatePost struct {
	Id      string `json:"id"`
	User_id string `json:"user_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type CreatedPost struct {
	Link string `json:"link"`
}

type List struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type ListRes struct {
	Posts Post  `json:"posts"`
	Count int64 `json:"count"`
}
