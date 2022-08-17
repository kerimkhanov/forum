package service

import (
	"fmt"
	"time"

	"forum/internal/models"

	uuid "github.com/satori/go.uuid"
)

func (s *AuthService) GetUserWithSession(session string) (models.Users, error) {
	user, err := s.storage.GetUserWithSession(session)
	if err != nil {
		return models.Users{}, fmt.Errorf("service.ParseSession - GetUser: %v", err)
	}
	return user, nil
}

func (s *AuthService) SetSession(username, password string) (models.Users, error) {
	user, err := s.storage.GetUserWithoutSession(username)
	if err != nil {
		return models.Users{}, fmt.Errorf("service.SetSession - GetUser: %v", err)
	}
	if err := CheckPasswordHash(password, user.Password); err != nil {
		return models.Users{}, fmt.Errorf("service.SetSession - CheckPasswwordHash: %v", err)
	}
	user.Session_token = uuid.NewV4().String()
	user.TimeSessions = time.Now().Add(10 * time.Hour)

	if err := s.storage.CreateSession(username, user); err != nil {
		return models.Users{}, fmt.Errorf("service.SetSession - CreateSession: %v", err)
	}
	return user, nil
}
