package storage

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
)

type Database struct {
	Auth
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		Auth: newStorage(db),
	}
}

type Auth interface {
	CreateUsers(user models.Users) error
	GetUsers() ([]models.Users, error)
}

type AuthStorage struct {
	db *sql.DB
}

func newStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (d *AuthStorage) CreateUsers(user models.Users) error {
	records := `INSERT INTO users(Login, Email, Password) values(?, ?, ?)`
	query, err := d.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(user.Login, user.Email, user.Password)
	if err != nil {
		return err
	}
	fmt.Print(user)
	return nil
}

func (d *AuthStorage) GetUsers(email) ([]models.Users, error) {
	var users []models.Users
	records := `SELCT * FROM users WHERE email = %s` 
	rows, err := d.db.Query(records)
	if err != nil {
		return users, err
	}
	var user models.Users
	for rows.Next() {
		err = rows.Scan(&user.Login, &user.Email, &user.Password)
		if err != nil {
			return users, err
		}
		users = append(users, user)
		// fmt.Println(err)
	}
	// fmt.Println(user.Email)
	rows.Close()
	return users, nil
}
