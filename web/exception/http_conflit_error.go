package exception

import "net/http"

type HttpConflictError struct {
	message string
	status  int
}

func NewHttpConflictError(message string) *HttpConflictError {
	return &HttpConflictError{
		message: message,
		status:  http.StatusConflict,
	}
}

func (e *HttpConflictError) Error() string {
	return e.message
}

func (e *HttpConflictError) Status() int {
	return e.status
}
