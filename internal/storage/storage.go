package storage

import (
	"database/sql"
)

type Database struct {
	Auth
	Post
	Like
	LikeDislikePost
	LikeDislikeComment
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		Auth:               newAuthStorage(db),
		Post:               newPostStorage(db),
		Like:               newLikeStorage(db),
		LikeDislikePost:    newLikeDislikePostStorage(db),
		LikeDislikeComment: newLikeDislikeCommentStorage(db),
	}
}
