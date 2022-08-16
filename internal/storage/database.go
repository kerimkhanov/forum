package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDb() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		os.Exit(1)
	}
	createTable(db)
	return db
}

func createTable(db *sql.DB) {
	userTable := `CREATE TABLE IF NOT EXISTS Users 
	(id INTEGER PRIMARY KEY NOT NULL,
	Login VARCHAR(64) NOT NULL,
	Email VARCHAR(64) NOT NULL,
	Password VARCHAR(64) NOT NULL,
	Session_token VARCHAR(64),
	TimeSessions DATE);`
	postTable := `CREATE TABLE IF NOT EXISTS post (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title VARCHAR(64) NOT NULL,
	text TEXT NOT NULL,
	author VARCHAR(64) NOT NULL);`
	err := dbExec(db, userTable, postTable)
	if err != nil {
		fmt.Errorf("error storage - database.go - dbExec")
	}
	fmt.Println("Table created successfully!")
}

func dbExec(db *sql.DB, query ...string) error {
	for i := 0; i < len(query); i++ {
		_, err := db.Exec(query[i])
		if err != nil {
			return fmt.Errorf("error delivery - db.go - dbExec %s", err)
		}
	}
	return nil
}

func AddUsers(db *sql.DB, Login string, Email string, Password string) {
	records := `INSERT INTO users(Login, Email, Password) VALUES (?, ?, ?)`
	query, err := db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(Login, Email, Password)
	if err != nil {
		log.Fatal(err)
	}
}
