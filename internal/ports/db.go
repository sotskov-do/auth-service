package ports

import (
	"context"

	"github.com/sotskov-do/auth-service/internal/domain/auth/models"
)

type Db interface {
	Login(ctx context.Context, user *models.User) (string, error)
	Register(ctx context.Context, user *models.User) error
}
