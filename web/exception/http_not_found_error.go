package exception

import "net/http"

type HttpNotFoundError struct {
	message string
	status  int
}

func NewHttpNotFoundError(message string) *HttpNotFoundError {
	return &HttpNotFoundError{
		message: message,
		status:  http.StatusNotFound,
	}
}

func (e *HttpNotFoundError) Error() string {
	return e.message
}

func (e *HttpNotFoundError) Status() int {
	return e.status
}
