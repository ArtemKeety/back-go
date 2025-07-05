package repository

import (
	"context"
	"database/sql"
	"github.com/ArtemKeety/back-go.git/internal/model"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (r *SessionRepository) AddSession(ctx context.Context, s model.Session) error {
	query := `INSERT INTO session (refresh, user_guid, time, ip) VALUES ($1, $2, $3, $4)`
	if _, err := r.db.ExecContext(ctx, query, s.Refresh, s.Guid, s.Time, s.Ip); err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) Update(ctx context.Context, s model.Session) error {
	query := `UPDATE session SET refresh = $1, time = $2 WHERE id = $3`
	if _, err := r.db.ExecContext(ctx, query, s.Refresh, s.Time, s.Id); err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) GetByToken(ctx context.Context, token string) (model.Session, error) {
	var s model.Session
	query := `SELECT s.id, s.refresh, s.user_guid, s.time, s.ip FROM session s WHERE s.refresh = $1 `
	res := r.db.QueryRowContext(ctx, query, token)
	if err := res.Scan(&s.Id, &s.Refresh, &s.Guid, &s.Time, &s.Ip); err != nil {
		return s, err
	}
	return s, nil
}
