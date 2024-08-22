package service

import (
	"BookHub/internal/repository"
)

type Service struct {
	IAuthorizationService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		IAuthorizationService: NewAuthService(repos.IAuthorizationRepo),
	}
}
