package model

type ErrorResponse struct {
	Message string `json:"message"`
	Err     error  `json:"error"`
}
