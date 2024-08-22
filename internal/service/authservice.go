package service

import (
	"BookHub/internal/model"
	"BookHub/internal/repository"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt      = "asdasdasd"
	tokenTTL  = 12 * time.Hour
	signinKey = "@#1sdafsadf"
)

type IAuthorizationService interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(login, password string) (string, error)
}

type AuthService struct {
	repo repository.IAuthorizationRepo
}

func NewAuthService(repo repository.IAuthorizationRepo) *AuthService {
	return &AuthService{repo: repo}
}

type tokenClaims struct {
	jwt.StandardClaims
	Userid int `json:"userid"`
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.PasswordHash = generatePasswordHash(user.PasswordHash)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signinKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
