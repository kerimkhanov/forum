package service

import (
	"fmt"
	"strconv"
	"time"

	"forum/internal/models"
	"forum/internal/storage"

	uuid "github.com/satori/go.uuid"
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
	CreateSession(userid int) (models.Users, error)
	DeleteUserSession(session string) error
	GetUserWithSession(session string) (models.Users, error)
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

func (s *AuthService) CreateSession(userid int) (models.Users, error) {
	user, err := s.storage.GetUserWithoutSession(username)
	if err != nil {
		return models.Users{}, fmt.Errorf("service.SetSession - GetUser: %v", err)
	}
	if err := CheckPasswordHash(password, user.Password); err != nil {
		return models.Users{}, fmt.Errorf("service.SetSession - CheckPasswwordHash: %v", err)
	}
	user.Session_token = uuid.

	var user models.Users
	user.Id = strconv.Itoa(userid)
	user.Session_token = uuid.NewV4().String()
	user.TimeSessions = time.Now()
	s.storage.CreateSession(user)
	return user, nil
}

func (s *AuthService) DeleteUserSession(session string) error {
	err := s.storage.DeleteUserSession(session)
	if err != nil {
		return fmt.Errorf("service.ParseSession - GetUser: %v", err)
	}
	return nil
}
