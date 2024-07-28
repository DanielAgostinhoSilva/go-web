package exception

import "net/http"

type HttpInternalServerError struct {
	message string
	status  int
}

func NewHttpInternalServerError(message string) *HttpInternalServerError {
	return &HttpInternalServerError{
		message: message,
		status:  http.StatusInternalServerError,
	}
}

func (e *HttpInternalServerError) Error() string {
	return e.message
}

func (e *HttpInternalServerError) Status() int {
	return e.status
}
