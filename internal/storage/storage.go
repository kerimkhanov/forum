package storage

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
	"log"
	"time"
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
	CreateSession(userid int, uuid string, sessionTime time.Time)
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
	return nil
}

func (d *AuthStorage) GetUser(email string) (models.Users, error) {
	var user models.Users
	records := fmt.Sprintf(`SELECT * FROM users WHERE Email = "%s"`, email)
	if err := d.db.QueryRow(records).Scan(&user.Id, &user.Login, &user.Email, &user.Password); err != nil {
		if err != nil {
			fmt.Println("GetUser", err)
			return user, fmt.Errorf("canPurchase %d: unknown album", user)
		}
	}
	return user, nil
}

func (d *AuthStorage) CreateSession(userid int, uuid string, dateTime time.Time) {
	InsertRequest := `INSERT INTO Sessions (Value, UserId, TimeSessions) values (? , ? , ?)`
	query, err := d.db.Prepare(InsertRequest)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(uuid, userid, dateTime.Add(120*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
}
