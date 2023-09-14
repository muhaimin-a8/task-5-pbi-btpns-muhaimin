package model

type WebResponse struct {
	Status  Status      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Status string

const (
	// Success when request successfully
	Success Status = "success"

	// Fail when user request because client error
	Fail Status = "fail"

	// Error when internal server error
	Error Status = "error"
)
