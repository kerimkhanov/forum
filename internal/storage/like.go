package storage

import (
	"database/sql"
)

type Like interface{}

type LikeStorage struct {
	db *sql.DB
}

func newLikeStorage(db *sql.DB) *LikeStorage {
	return &LikeStorage{
		db: db,
	}
}
