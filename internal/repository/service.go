package repository

import (
	"fmt"
	"forum/internal/models"
	"forum/internal/storage"
	"time"
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
	UserByEmail(email string) (models.Users, error)
	CreateSession(userid int, uuid string, sessionTime time.Time)
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

func (s *AuthService) UserByEmail(email string) (models.Users, error) {
	// user.Password = hash
	user, err := s.storage.GetUser(email)
	if err != nil {
		fmt.Errorf("Service -> UserByEmail: %v", err)
	}
	return user, err
}

func (s *AuthService) CreateSession(userid int, uuid string, datetime time.Time) {
	s.storage.CreateSession(userid, uuid, datetime)
}
