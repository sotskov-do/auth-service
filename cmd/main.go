package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sotskov-do/auth-service/internal/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	go application.Start(ctx)
	<-ctx.Done()
	application.Stop(ctx)
}
