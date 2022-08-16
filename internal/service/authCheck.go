package service

import (
	"errors"

	"forum/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

// check password in hash
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func CorrectAuth(user models.Users, Email, Password string) error {
	if Email != user.Email {
		return errors.New("Invalid email address")
	}
	if !CheckPasswordHash(Password, user.Password) {
		return errors.New("password incorrect")
	}
	return nil
}
