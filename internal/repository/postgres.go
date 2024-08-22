package repository

import (
	"BookHub/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewConnectionDB(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Db.Host, cfg.Db.Port, cfg.Db.Username, cfg.Db.Dbname, cfg.Db.Password, cfg.Db.SslMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
