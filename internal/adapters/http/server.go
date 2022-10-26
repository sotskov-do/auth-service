package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sotskov-do/auth-service/internal/ports"
	"github.com/sotskov-do/auth-service/pkg/config"
	// httpMiddleware "gitlab.com/g6834/team26/task/pkg/middleware"
)

type Server struct {
	task     ports.Auth
	server   *http.Server
	listener net.Listener
	config   *config.Config
	port     int
}

func New(task ports.Auth, config *config.Config) (*Server, error) {
	var (
		err error
		s   Server
	)
	port := fmt.Sprintf(":%s", config.Server.Port)
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed listen port", err)
	}
	s.config = config
	s.task = task
	s.port = s.listener.Addr().(*net.TCPAddr).Port

	s.server = &http.Server{
		Handler: s.routes(),
	}

	return &s, nil
}

func (s *Server) Port() int {
	return s.port
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.server.Serve(s.listener); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) routes() http.Handler {
	r := chi.NewRouter()
	r.Mount("/auth/v1", s.authHandlers())

	return r
}
