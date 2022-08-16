package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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
	GetUser(email string) (models.Users, error)
	CreateSession(user models.Users)
	GetUserWithoutSession(login string) (models.Users, error)
	GetUserWithSession(token string) (models.Users, error)
	DeleteUserSession(token string) error
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
	_, err = query.Exec(&user.Login, &user.Email, &user.Password)
	if err != nil {
		fmt.Printf("\n %s Login = %s Email = %s Password = %s\n", "Create Users", user.Login, user.Email, user.Password)
		return err
	}
	return nil
}

func (d *AuthStorage) GetUser(email string) (models.Users, error) {
	var user models.Users
	// records := fmt.Sprintf(`SELECT id, Login, Email, Password FROM Users WHERE Email = "%s"`, email)
	fmt.Println(email)
	records := fmt.Sprintf(`SELECT	*
	FROM Users
	WHERE Email = "%s";`, email)
	query := d.db.QueryRow(records)
	// fmt.Println(query.Columns())
	if err := query.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.Session_token, &user.TimeSessions); err != nil {
		fmt.Println("GetUser", user.Id)
		fmt.Println("GetUser", user.Login)
		fmt.Println("GetUser", user.Email)
		fmt.Println("GetUser", user.Password)
		fmt.Println("GetUser", user.Session_token)
		fmt.Println("GetUser", user.TimeSessions)
		fmt.Println("GetUser", err)
		return user, fmt.Errorf("canPurchase %d: unknown album", user)
	}
	return user, nil
}

func (d *AuthStorage) CreateSession(user models.Users) {
	InsertRequest := `UPDATE Sessions (Value, UserId, TimeSessions) values (? , ? , ?)`
	query, err := d.db.Prepare(InsertRequest)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(user.Session_token, user.Id, time.Now().Add(120*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
}
