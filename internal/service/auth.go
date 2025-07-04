package service

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/ArtemKeety/back-go.git/internal/model"
	"github.com/ArtemKeety/back-go.git/internal/repository"
	"github.com/ArtemKeety/back-go.git/pkg/hashing"
	"github.com/ArtemKeety/back-go.git/pkg/token"
	"time"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, u model.UserRequest) (string, error) {

	ex, err := s.repo.Auth.CheckUserExists(ctx, u)
	if err != nil {
		return "", err
	}

	if ex.Guid != "" {
		return "", errors.New("user already exists")
	}

	if u.Password, err = hashing.HashPassword(u.Password); err != nil {
		return "", errors.New("error hashing password")
	}

	guid, err := s.repo.Auth.AddUser(ctx, u)
	if err != nil {
		return "", err
	}

	return guid, nil
}

func (s *AuthService) Login(ctx context.Context, ip string, u model.UserRequest) (map[string]string, error) {
	user, err := s.repo.Auth.CheckUserExists(ctx, u)
	if err != nil {
		return nil, err
	}

	if user.Guid == "" {
		return nil, errors.New("user does not exist")
	}

	if !hashing.CheckPasswordHash(u.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	data := make(map[string]string)

	refreshToken, err := token.NewRefreshToken(user.Guid)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Session.AddSession(
		ctx,
		model.Session{
			Refresh: refreshToken,
			Guid:    user.Guid,
			Time:    time.Now().Add(token.RefreshTime * time.Hour).UTC(),
			Ip:      ip,
		}); err != nil {
		return nil, err
	}

	accessToken, err := token.NewAccessToken(user.Guid)
	if err != nil {
		return nil, err
	}

	data["access_token"] = accessToken
	data["refresh_token"] = base64.StdEncoding.EncodeToString([]byte(refreshToken))

	return data, nil
}
