package service

import (
	"context"
	"errors"
	"github.com/ArtemKeety/back-go.git/internal/model"
	"github.com/ArtemKeety/back-go.git/internal/repository"
	"github.com/ArtemKeety/back-go.git/pkg/hashing"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, u model.UserRequest) (string, error) {

	ex, err := s.repo.CheckUserExists(ctx, u)
	if err != nil {
		return "", err
	}

	if ex.Guid != "" {
		return "", errors.New("user already exists")
	}

	if u.Password, err = hashing.HashPassword(u.Password); err != nil {
		return "", errors.New("error hashing password")
	}

	guid, err := s.repo.AddUser(ctx, u)
	if err != nil {
		return "", err
	}

	return guid, nil
}
