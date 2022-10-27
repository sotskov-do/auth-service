package application

import (
	"context"
	"log"
	"os"

	"github.com/sotskov-do/auth-service/internal/adapters/http"
	"github.com/sotskov-do/auth-service/internal/adapters/postgres"
	"github.com/sotskov-do/auth-service/internal/domain/auth"
	"github.com/sotskov-do/auth-service/pkg/config"
	"golang.org/x/sync/errgroup"
)

var (
	s           *http.Server
	authService *auth.Service
	db          *postgres.PostgresDatabase
)

func Start(ctx context.Context) {
	c, err := config.New()
	if err != nil {
		log.Fatalf("Error parsing env: %s", err)
	}

	db, err = postgres.New(ctx, c.Server.PgUrl)
	if err != nil {
		log.Fatalf("db init failed: %s", err)
		os.Exit(1)
	}

	authService = auth.New(db, c)

	s, err = http.New(authService, c)
	if err != nil {
		log.Fatalf("http server creating failed: %s", err)
		os.Exit(1)
	}

	var g errgroup.Group
	g.Go(func() error {
		return s.Start(ctx)
	})

	log.Println("app is started")
	err = g.Wait()
	if err != nil {
		log.Fatalf("http server start failed: %s", err)
		os.Exit(1)
	}
}

func Stop(ctx context.Context) {
	_ = authService.Stop()
	_ = s.Stop(ctx)
	_ = db.Stop(ctx)
	log.Println("app has stopped")
}
