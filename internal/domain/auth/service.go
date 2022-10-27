package auth

import (
	"context"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/sotskov-do/auth-service/internal/domain/auth/models"
	"github.com/sotskov-do/auth-service/internal/ports"
	"github.com/sotskov-do/auth-service/pkg/config"
)

type Service struct {
	db    ports.Db
	token *models.TokenAuth
}

func New(db ports.Db, c *config.Config) *Service {
	return &Service{
		db:    db,
		token: &models.TokenAuth{Token: jwtauth.New("HS256", []byte(c.Server.SecretKey), nil)},
	}
}

func (s *Service) Stop() error {
	return nil
}

func (s *Service) Login(ctx context.Context, user *models.User) (string, error) {
	password, err := s.db.Login(ctx, user)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (s *Service) Register(ctx context.Context, user *models.User) error {
	err := s.db.Register(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GenToken(ctx context.Context, user *models.User) (string, error) {
	tokenClaim := map[string]interface{}{}
	tokenClaim["login"] = user.Login
	tokenClaim["expires"] = time.Now().Add(time.Hour)
	_, token, err := s.token.Token.Encode(tokenClaim)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Service) GetToken() *models.TokenAuth {
	return s.token
}
