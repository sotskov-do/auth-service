package auth

import (
	"context"

	"github.com/sotskov-do/auth-service/internal/ports"
)

type Service struct {
	db ports.Db
}

func New(db ports.Db) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Stop() error {
	return nil
}

func (s *Service) Login(ctx context.Context, login string) error {
	return nil
}

func (s *Service) Register(ctx context.Context, login string) error {
	return nil
}
