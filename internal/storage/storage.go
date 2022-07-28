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

func (d *AuthStorage) GetUsers(user models.Users) error {
	fmt.Print(user.Email)
	records := `SELCT * FROM users`
	rows, err := d.db.Query(records)
	if err != nil {
		return err
	}
	for rows.Next() {
		err = rows.Scan(&user.Login, &user.Email, &user.Password)
		// fmt.Println(err)
	}
	// fmt.Println(user.Email)
	rows.Close()
	return nil
}
