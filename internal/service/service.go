package service

import (
	"context"
	"github.com/ArtemKeety/back-go.git/internal/model"
	"github.com/ArtemKeety/back-go.git/internal/repository"
)

type Auth interface {
	CreateUser(ctx context.Context, u model.UserRequest) (string, error)
	Login(ctx context.Context, ip string, u model.UserRequest) (map[string]string, error)
}

type Session interface {
	ChangeToken(ctx context.Context, ip string, token string) (map[string]string, error)
	CloseSession(ctx context.Context, t string) error
}

type Service struct {
	Auth
	Session
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth:    NewAuthService(repo),
		Session: NewSessionService(repo),
	}
}
