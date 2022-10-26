package ports

import "context"

type Auth interface {
	Login(ctx context.Context, login string) error
	Register(ctx context.Context, login string) error
}
