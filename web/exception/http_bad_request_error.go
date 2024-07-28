package exception

import "net/http"

type HttpBadRequestError struct {
	message string
	status  int
}

func NewHttpBadRequestError(message string) *HttpBadRequestError {
	return &HttpBadRequestError{
		message: message,
		status: http.StatusBadRequest,
	}
}

func (e *HttpBadRequestError) Error() string {
	return e.message
}

func (e *HttpBadRequestError) Status() int {
	return e.status
}
