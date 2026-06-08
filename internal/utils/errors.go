package utils

import "fmt"

type AppError struct {
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("status %d: %s", e.StatusCode, e.Message)
}

func NewAppError(status int, msg string) *AppError {
	return &AppError{
		StatusCode: status,
		Message:    msg,
	}
}