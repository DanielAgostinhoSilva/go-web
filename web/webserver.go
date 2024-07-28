package web

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	Router        chi.Router
	Handlers      map[string]Handler
	Controllers   []Controller
	WebServerPort string
}

func NewServer(webServerPort string) *Server {
	return &Server{WebServerPort: webServerPort}
}

func (s *Server) AddController(controller Controller) {
	s.Controllers = append(s.Controllers, controller)
}

func (s *Server) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers[path] = Handler{Method: method, HandlerFunc: handler}
}

func (s *Server) AddMiddleware(middleware func(http.Handler) http.Handler) {
	s.Router.Use(middleware)
}

func (s *Server) Start() {
	for _, controller := range s.Controllers {
		s.Router.Route(controller.Path(), controller.Router())
	}
	for path, handler := range s.Handlers {
		s.Router.Method(handler.Method, path, handler.HandlerFunc)
	}
	http.ListenAndServe(":"+s.WebServerPort, s.Router)
}
