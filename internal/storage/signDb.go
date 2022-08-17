package storage

import (
	"fmt"
	"time"

	"forum/internal/models"
)

func (d *AuthStorage) GetUserWithoutSession(email string) (models.Users, error) {
	var user models.Users
	query := "Select login, email, password FROM users where email=$1"
	row := d.db.QueryRow(query, email)
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
	fmt.Printf("\n token --->  %s\n", token)
	query := "UPDATE users SET session_token=$1, TimeSessions=$2 WHERE session_token=$3"
	sqlResult, err := d.db.Exec(query, "", time.Now(), token)
	fmt.Printf("\nsql result = %s\n", sqlResult)
	if err != nil {
		fmt.Println("storage - signDB - getUserWithSession")
		return fmt.Errorf("Storage.GetUser: %v", err)
	}
	return nil
}
