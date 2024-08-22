package repository

import (
	"BookHub/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
)

type IAuthorizationRepo interface {
	CreateUser(user model.User) (int, error)
	GetUser(login string, password string) (model.User, error)
}

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPosgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("Insert into %s (Username, Email, Password) values ($1, $2, $3) Returning id", userTable)
	row := r.db.QueryRow(query, user.Username, user.Email, user.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login string, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("Select id from %s where Login=$1 AND password=$2 ", userTable)
	err := r.db.Get(&user, query, login, password)
	return user, err
}
