package web

import (
	"github.com/go-chi/chi/v5"
)

type Controller interface {
	Router() func(r chi.Router)
	Path() string
}
