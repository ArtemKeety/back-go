package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConfDb struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SslMode  string
}

func NewDB(cfg ConfDb) (*sql.DB, error) {
	str := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SslMode)

	db, err := sql.Open("postgres", str)
	if err != nil {
		return nil, err
	}

	return db, nil
}
