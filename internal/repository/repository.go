package repository

import (
	"context"
	"database/sql"
	"github.com/ArtemKeety/back-go.git/internal/model"
)

type Auth interface {
	AddUser(ctx context.Context, u model.UserRequest) (string, error)
	CheckUserExists(ctx context.Context, u model.UserRequest) (bool, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
	}
}
