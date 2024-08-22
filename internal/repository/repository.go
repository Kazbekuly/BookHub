package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	IAuthorizationRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IAuthorizationRepo: NewAuthPosgres(db),
	}
}
