package service

import (
	"context"
	"errors"
	"github.com/ArtemKeety/back-go.git/internal/model"
	"github.com/ArtemKeety/back-go.git/internal/repository"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, u model.UserRequest) (int, error) {

	userFlag, err := s.repo.CheckUserExists(ctx, u)
	if err != nil {
		return -1, err
	}

	if userFlag {
		return -1, errors.New("user already exists")
	}

	id, err := s.repo.AddUser(ctx, u)
	if err != nil {
		return -1, err
	}

	return id, nil
}
