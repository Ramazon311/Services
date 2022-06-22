package models

// Error ...
type Error struct {
	Message string `json:"message"`
}

// StandardErrorModel ...
type StandardErrorModel struct {
	Error Error `json:"error"`
}

// ResponseSuccess ...
type ResponseSuccess struct {
	Metadata interface{}
	Data     interface{}
}

// ResponseError ...
type ResponseError struct {
	Error interface{} `json:"error"`
}

// ServerError ...
type ServerError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ValidationError ...
type ValidationError struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	UserMessage string
}

//InternalServerError ...
type InternalServerError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
