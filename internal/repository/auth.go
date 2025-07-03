package repository

import (
	"context"
	"github.com/ArtemKeety/back-go.git/internal/model"
)

type AuthRepository struct {
	db interface{}
}

func NewAuthRepo(db interface{}) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a AuthRepository) AddUser(ctx context.Context, u model.UserRequest) (int, error) {

	return -1, nil
}

func (a AuthRepository) CheckUserExists(ctx context.Context, u model.UserRequest) (bool, error) {

	return false, nil
}
