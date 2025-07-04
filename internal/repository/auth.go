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

func (r *AuthRepository) AddUser(ctx context.Context, u model.UserRequest) (string, error) {
	query := "INSERT INTO users (guid, login, email, password) VALUES ($1, $2, $3, $4)"
	guid := uuid.New().String()

	_, err := r.db.ExecContext(ctx, query, guid, u.Login, u.Email, u.Password)
	if err != nil {
		return "", err
	}

	return guid, nil
}

func (r *AuthRepository) CheckUserExists(ctx context.Context, u model.UserRequest) (model.User, error) {
	var user model.User
	row := r.db.QueryRowContext(
		ctx, "SELECT guid, login, email, password FROM users WHERE email = $1 OR login = $2", u.Email, u.Login)
	if err := row.Scan(&user.Guid, &user.Login, &user.Email, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, nil
		}
		return user, err
	}
	return user, nil
}
