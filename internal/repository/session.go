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
