package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/database.db")
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
 TimeSessions DATE,
 posts INT DEFAULT 0,
 likes INT DEFAULT 0,
 dislikes INT DEFAULT 0,
 comments INT DEFAULT 0);`
	postTable := `CREATE TABLE IF NOT EXISTS post (
 id INTEGER PRIMARY KEY AUTOINCREMENT,
 title VARCHAR(64) NOT NULL,
 text TEXT NOT NULL,
 author VARCHAR(64) NOT NULL,
 created_at DATE DEFAULT (datetime('now','localtime')),
 likes INT DEFAULT 0,
 dislikes INT DEFAULT 0,
 comments INT DEFAULT 0);`
	tagsTable := `CREATE TABLE IF NOT EXISTS tags (
 post_id INTEGER,
 tag TEXT,  
 FOREIGN KEY (post_id) REFERENCES post(id));`
	commentsTable := `CREATE TABLE IF NOT EXISTS comments (
  comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
  post_id INTEGER,
  comment TEXT NOT NULL,
  login VARCHAR(64) NOT NULL,
  likes INT DEFAULT 0,
  dislikes INT DEFAULT 0,
  FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE);`
	likeTable := `CREATE TABLE IF NOT EXISTS likes (
  post_id INTEGER,
  comment_id INTEGER,
  login TEXT NOT NULL,
  FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
  FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE
  );`
	dislikeTable := `CREATE TABLE IF NOT EXISTS dislikes (
  post_id INTEGER,
  comment_id INTEGER,
  login TEXT NOT NULL,
  FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
  FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE
  );`

	err := dbExec(db, userTable, postTable, tagsTable, commentsTable, likeTable, dislikeTable)
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
