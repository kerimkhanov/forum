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
		fmt.Println(err)
		fmt.Println("asd")
		os.Exit(1)
	}
	createTable(db)
	return db
}

func createTable(db *sql.DB) {
	CREATE_TABLE := `CREATE TABLE IF NOT EXISTS Users 
	(id INTEGER PRIMARY KEY NOT NULL,
	Login VARCHAR(64) NOT NULL,
	Email VARCHAR(64) NOT NULL,
	Password VARCHAR(64) NOT NULL);
	CREATE TABLE IF NOT EXISTS Sessions
	(id INTEGER PRIMARY KEY NOT NULL,
	Value VARCHAR(64) NOT NULL,
	UserId INTEGER,
	TimeSessions DATE,
	FOREIGN KEY(UserId) REFERENCES Users(id));`
	_, err := db.Exec(CREATE_TABLE)
	if err != nil {
		fmt.Println("asd")
		fmt.Println(err)
	}
	fmt.Println("Table created successfully!")
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
