package storage

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
)

type Auth interface {
	CreateUsers(user models.Users) error
	GetUser(email string) (models.Users, error)
	CreateSession(email string, user models.Users) error
	GetUserWithoutSession(login string) (models.Users, error)
	GetUserWithSession(token string) (models.Users, error)
	DeleteUserSession(token string) error
}

type AuthStorage struct {
	db *sql.DB
}

func newAuthStorage(db *sql.DB) *AuthStorage {
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
	_, err = query.Exec(&user.Login, &user.Email, &user.Password)
	if err != nil {
		return fmt.Errorf("\n %s Login = %s Email = %s Password = %s : -- %v\n", "Create Users", user.Login, user.Email, user.Password, err)
	}
	return nil
}

func (d *AuthStorage) GetUser(email string) (models.Users, error) {
	var user models.Users
	records := fmt.Sprintf(`SELECT id, Login, Email, Password FROM Users WHERE Email = "%s"`, email)
	fmt.Println(email)
	// records := fmt.Sprintf(`SELECT	*
	// FROM Users
	// WHERE Email = "%s";`, email)
	query := d.db.QueryRow(records)
	// fmt.Println(query.Columns())
	if err := query.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.Session_token, &user.TimeSessions); err != nil {
		return user, fmt.Errorf("canPurchase %d: unknown album", user)
	}
	return user, nil
}

func (d *AuthStorage) CreateSession(email string, user models.Users) error {
	InsertRequest := `UPDATE users SET Session_token=$1, TimeSessions=$2 WHERE email=$3`
	_, err := d.db.Exec(InsertRequest, user.Session_token, user.TimeSessions, email)
	if err != nil {
		return fmt.Errorf("storage - CreateSession %v", err)
	}
	return nil
}
