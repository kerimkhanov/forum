package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDb() *sql.DB {
	// f, err := os.OpenFile("database.db", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	// if err != nil {
	// 	log.Fatal(err)
	// }
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
	CREATE_TABLE := "CREATE TABLE IF NOT EXISTS Users" +
		"(id INTEGER PRIMARY KEY NOT NULL," +
		"Login VARCHAR(64) NOT NULL," +
		"Email VARCHAR(64) NOT NULL," +
		"Password VARCHAR(64) NOT NULL)"
	// statement, err := db.Prepare(CREATE_TABLE)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	_, err := db.Exec(CREATE_TABLE)
	if err != nil {
		fmt.Println("asd")
		fmt.Println(err)
	}

	// query, err := db.Prepare(statement.Close().Error())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// query.Exec()
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
