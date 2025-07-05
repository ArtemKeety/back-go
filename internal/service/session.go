package service

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/ArtemKeety/back-go.git/internal/repository"
	"github.com/ArtemKeety/back-go.git/pkg/token"
	"github.com/sirupsen/logrus"
	"time"
)

type SessionService struct {
	repo *repository.Repository
}

func NewSessionService(repo *repository.Repository) *SessionService {
	return &SessionService{
		repo: repo,
	}
}

func (s *SessionService) ChangeToken(ctx context.Context, ip string, t string) (map[string]string, error) {

	session, err := s.repo.Session.GetByToken(ctx, t)
	if err != nil {
		return nil, errors.New("session doesn't exist")
	}

	if session.Ip != ip {
		logrus.Infof("send webhok %s", session.Ip)
	}

	durationTime := session.Time.Sub(time.Now()).Seconds()
	if durationTime < 0 {
		if err := s.repo.Session.DeleteByToken(ctx, t); err != nil {
			logrus.Errorf("delete session fail %s", err.Error())
		}
		return nil, errors.New("session timeout")
	}

	data := make(map[string]string)

	refreshToken, err := token.NewRefreshToken(session.Guid)
	if err != nil {
		return nil, err
	}

	accessToken, err := token.NewAccessToken(session.Guid)
	if err != nil {
		return nil, err
	}

	session.Time = time.Now().Add(token.RefreshTime * time.Hour).UTC()
	session.Refresh = refreshToken

	data["access_token"] = accessToken
	data["refresh_token"] = base64.StdEncoding.EncodeToString([]byte(refreshToken))

	if err := s.repo.Session.Update(ctx, session); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *SessionService) CloseSession(ctx context.Context, t string) error {
	return s.repo.Session.DeleteByToken(ctx, t)
}
