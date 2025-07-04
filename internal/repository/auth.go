package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ArtemKeety/back-go.git/internal/model"
	"github.com/google/uuid"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a AuthRepository) AddUser(ctx context.Context, u model.UserRequest) (string, error) {
	query := "INSERT INTO users (guid, login, email, password) VALUES ($1, $2, $3, $4)"
	guid := uuid.New().String()

	_, err := a.db.ExecContext(ctx, query, guid, u.Login, u.Email, u.Password)
	if err != nil {
		return "", err
	}

	return guid, nil
}

func (a AuthRepository) CheckUserExists(ctx context.Context, u model.UserRequest) (bool, error) {
	var guid string
	row := a.db.QueryRowContext(ctx, "SELECT guid FROM users WHERE email = $1 OR login = $2", u.Email, u.Login)
	if err := row.Scan(&guid); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
