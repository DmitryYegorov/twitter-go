package utils

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewHttpError(status int, message string) *HttpError {
	return &HttpError{status, message}
}
