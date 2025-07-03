package service

import (
	"context"
	"github.com/ArtemKeety/back-go.git/internal/model"
	"github.com/ArtemKeety/back-go.git/internal/repository"
)

type Auth interface {
	CreateUser(ctx context.Context, u model.UserRequest) (int, error)
}

type Service struct {
	Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
