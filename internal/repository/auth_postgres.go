package repository

import (
	"BookHub/internal/model"
	"fmt"
	"time"

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
	query := fmt.Sprintf("INSERT INTO %s (username, email, password, createdAt, updatedAt) values ($1, $2, $3, $4, $5) RETURNING id", userTable)
	createdAt := time.Now()
	row := r.db.QueryRow(query, user.Username, user.Email, user.PasswordHash, createdAt, createdAt)
	if err := row.Scan(&id); err != nil {
		fmt.Println(err)
		return 0, nil
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username string, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("Select id from %s where username=$1 AND password=$2 ", userTable)
	err := r.db.Get(&user, query, username, password)
	fmt.Println(err)
	return user, err
}
