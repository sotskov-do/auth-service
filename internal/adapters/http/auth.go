package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) authHandlers() http.Handler {
	r := chi.NewRouter()

	r.Post("/login", s.LoginHandler)
	r.Post("/register", s.RegisterHandler)

	return r
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("register"))
}
