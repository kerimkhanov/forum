package models

import "time"

type Posts struct {
	Id        string
	Title     string
	Text      string
	Author    string
	Tags      []string
	CreatedAt time.Time
	Likes     int
	Dislikes  int
	Comments  []Comments
}

type Allpost struct {
	Users
	Posts []Posts
}

type OnePost struct {
	Users
	Posts
}
