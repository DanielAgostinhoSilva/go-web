package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/DanielAgostinhoSilva/go-web/web/exception"
	"github.com/go-errors/errors"
	"net/http"
	"time"
)

type Problem struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
}

func NewProblemNotFound(detail string) *Problem {
	return &Problem{
		Status:    http.StatusNotFound,
		Timestamp: time.Now().Local(),
		Title:     "Resource not found",
		Detail:    detail,
	}
}

func NewProblemConflict(detail string) *Problem {
	return &Problem{
		Status:    http.StatusConflict,
		Timestamp: time.Now().Local(),
		Title:     "Resource is already in use",
		Detail:    detail,
	}
}

func NewProblemBadRequest(detail string) *Problem {
	return &Problem{
		Status:    http.StatusBadRequest,
		Timestamp: time.Now().Local(),
		Title:     "Invalid data",
		Detail:    detail,
	}
}

func NewProblemInternalServerError(detail string) *Problem {
	return &Problem{
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now().Local(),
		Title:     "Internal Server Error",
		Detail:    "An unexpected internal error has occurred in the system. Try again and if the problem persists, contact your system administrator.",
	}
}

func ExceptionHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer handler(w, r)
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err, ok := recover().(error); ok {
		var httpNotFoundError *exception.HttpBadRequestError
		var httpConflictError *exception.HttpConflictError
		var httpBadRequestError *exception.HttpBadRequestError
		var httpInternalServerError *exception.HttpInternalServerError
		switch {
		case errors.As(err, &httpNotFoundError):
			handlerError(NewProblemNotFound(err.Error()), w)
		case errors.As(err, &httpConflictError):
			handlerError(NewProblemConflict(err.Error()), w)
		case errors.As(err, &httpBadRequestError):
			handlerError(NewProblemBadRequest(err.Error()), w)
		case errors.As(err, &httpInternalServerError):
			handlerError(NewProblemInternalServerError(err.Error()), w)
		default:
			handlerError(NewProblemInternalServerError("An unexpected internal error has occurred in the system. Try again and if the problem persists, contact your system administrator."), w)
		}
		logError(err)
	}
}

func handlerError(problem *Problem, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(problem.Status)
	json.NewEncoder(w).Encode(problem)
}

func logError(err interface{}) {
	wrappedErr := errors.Wrap(err, 0)
	fmt.Println(wrappedErr.ErrorStack())
}
