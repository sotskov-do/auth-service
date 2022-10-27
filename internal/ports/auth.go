package ports

import (
	"context"

	"github.com/sotskov-do/auth-service/internal/domain/auth/models"
)

type Auth interface {
	Login(ctx context.Context, user *models.User) (string, error)
	Register(ctx context.Context, user *models.User) error
	GenToken(ctx context.Context, user *models.User) (string, error)
	GetToken() *models.TokenAuth
}
