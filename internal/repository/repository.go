package repository

import (
	"context"
	"github.com/ArtemKeety/back-go.git/internal/model"
)

type Auth interface {
	AddUser(ctx context.Context, u model.UserRequest) (int, error)
	CheckUserExists(ctx context.Context, u model.UserRequest) (bool, error)
}

type Repository struct {
	Auth
}

func NewRepository(db interface{}) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
	}
}
