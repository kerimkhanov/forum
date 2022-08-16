package storage

import (
	"fmt"
	"time"

	"forum/internal/models"
)

func (d *AuthStorage) GetUserWithoutSession(login string) (models.Users, error) {
	var user models.Users

	query := "Select login, email, password FROM users where login=$1"
	row := d.db.QueryRow(query, login)
	err := row.Scan(&user.Email, &user.Login, &user.Password)
	if err != nil {
		return models.Users{}, fmt.Errorf("storage.GetUser: %v", err)
	}
	return user, err
}

func (d *AuthStorage) GetUserWithSession(token string) (models.Users, error) {
	var user models.Users
	query := "Select login, email, password, Session_token, TimeSessions FROM users WHERE session_token=$1"
	row := d.db.QueryRow(query, token)
	err := row.Scan(&user.Login, &user.Email, &user.Password, &user.Session_token, &user.TimeSessions)
	if err != nil {
		return models.Users{}, fmt.Errorf("storage.GetUser: %v", err)
	}
	return user, err
}

func (d *AuthStorage) DeleteUserSession(token string) error {
	query := "DELETE user SET session_token=$1, TimeSessions=$2 WHERE session_token=$3"
	_, err := d.db.Exec(query, "", time.Now(), token)
	if err != nil {
		return fmt.Errorf("Storage.GetUser: %v", err)
	}
	return nil
}
