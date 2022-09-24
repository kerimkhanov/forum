package models

import "time"

type Users struct {
	Id              string
	Email           string
	Login           string
	Password        string
	CountOfPosts    int
	CountOfLikes    int
	CountOfDislikes int
	CountOfComments int
	Session_token   string
	TimeSessions    time.Time
}
