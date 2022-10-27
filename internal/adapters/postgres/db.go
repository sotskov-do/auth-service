package postgres

import (
	"context"
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
	"github.com/sotskov-do/auth-service/internal/domain/auth/models"
)

type PostgresDatabase struct {
	psqlClient *sql.DB
}

func New(ctx context.Context, pgconn string) (*PostgresDatabase, error) {
	db, err := sql.Open("postgres", pgconn+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &PostgresDatabase{psqlClient: db}, nil
}

func (pdb *PostgresDatabase) Stop(ctx context.Context) error {
	err := pdb.psqlClient.Close()
	if err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDatabase) Login(ctx context.Context, user *models.User) (string, error) {
	var password string
	query := `SELECT "password" FROM "auth" WHERE "login" = $1`
	row := pdb.psqlClient.QueryRowContext(ctx, query, user.Login)
	row.Scan(&password)
	password = strings.TrimSpace(password)
	return password, nil
}

func (pdb *PostgresDatabase) Register(ctx context.Context, user *models.User) error {
	query := `INSERT INTO "auth" ("login", "email", "password", "phone") VALUES	($1, $2, $3, $4)`
	_, err := pdb.psqlClient.ExecContext(ctx, query, user.Login, user.Email, user.EncodePassword(), user.Phone)
	if err != nil {
		return err
	}
	return nil
}
