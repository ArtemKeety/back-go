package repository

import (
	"context"
	"database/sql"
	"github.com/ArtemKeety/back-go.git/internal/model"
)

type Auth interface {
	AddUser(ctx context.Context, u model.UserRequest) (string, error)
	CheckUserExists(ctx context.Context, u model.UserRequest) (model.User, error)
}

type Session interface {
	AddSession(ctx context.Context, s model.Session) error
	GetByToken(ctx context.Context, t string) (model.Session, error)
	Update(ctx context.Context, s model.Session) error
}

type Repository struct {
	Auth
	Session
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Auth:    NewAuthRepo(db),
		Session: NewSessionRepository(db),
	}
}
