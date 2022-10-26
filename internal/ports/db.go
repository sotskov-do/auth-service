package ports

import "context"

type Db interface {
	Login(ctx context.Context) error
	Register(ctx context.Context) error
}
