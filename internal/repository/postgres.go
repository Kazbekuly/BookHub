package repository

import (
	"BookHub/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	err = CreateUserTable(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateUserTable(db *sqlx.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	sqlt, err := db.PrepareContext(ctx, `CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, 
		Username TEXT NOT NULL UNIQUE, 
		Email TEXT NOT NULL UNIQUE, 
		Password TEXT NOT NULL,
		CreatedAt TEXT NOT NULL,
		UpdatedAt TEXT NOT NULL)`)
	sqlt.Exec()
	if err != nil {
		return err
	}
	return nil
}
