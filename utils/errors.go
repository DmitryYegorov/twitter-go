package utils

import "fmt"

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (h *HttpError) Error() string {
	return fmt.Sprintf("Status %d, %s", h.Status, h.Message)
}

func NewHttpError(status int, message string) *HttpError {
	return &HttpError{status, message}
}
