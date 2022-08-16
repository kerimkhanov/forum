package service

import (
	"fmt"

	"forum/internal/models"
)

func (s *AuthService) GetUserWithSession(session string) (models.Users, error) {
	user, err := s.storage.GetUserWithSession(session)
	if err != nil {
		return models.Users{}, fmt.Errorf("service.ParseSession - GetUser: %v", err)
	}
	return user, nil
}
