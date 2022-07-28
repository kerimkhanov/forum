package internal

import (
	"forum/internal/models"
	"forum/internal/storage"
)

type Service struct {
	Auth
}

func NewService(s storage.Auth) *Service {
	return &Service{
		Auth: newAuthService(s),
	}
}

type Auth interface {
	AddUsers(user models.Users) error
	GetUsers(user models.Users) error
}

type AuthService struct {
	storage storage.Auth
}

func newAuthService(storage storage.Auth) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) AddUsers(user models.Users) error {
	// user.Password = hash
	return s.storage.CreateUsers(user)
}

func (s *AuthService) GetUsers(user models.Users) error {
	// user.Password = hash
	return s.storage.CreateUsers(user)
}
