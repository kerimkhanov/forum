package storage

import (
	"database/sql"
	"fmt"
)

type LikeDislikePost interface {
	AddPostLike(post_id int, login string) error
	AddPostDislike(post_id int, login string) error
	PostHasLike(post_id int, login string) error
	PostHasDislike(post_id int, login string) error
	RemoveLikeFromPost(post_id int, login string) error
	RemoveDislikeFromPost(post_id int, login string) error
	LikePost(post_id int, login string) error
	DislikePost(post_id int, login string) error
	GetPostLikes(post_id int) ([]string, error)
	GetPostDislikes(post_id int) ([]string, error)
}

type LikeDislikePostStorage struct {
	db *sql.DB
}

func newLikeDislikePostStorage(db *sql.DB) *LikeDislikePostStorage {
	return &LikeDislikePostStorage{
		db: db,
	}
}

func (d *LikeDislikePostStorage) GetPostDislikes(post_id int) ([]string, error) {
	var postLikes []string
	query := `SELECT login FROM dislikes WHERE post_id = $1`
	rows, err := d.db.Query(query, post_id)
	if err != nil {
		return nil, fmt.Errorf("storage.post-like-dislike.GetPostLikes.Query: %v", err)
	}
	for rows.Next() {
		var postLike string
		if err := rows.Scan(&postLike); err != nil {
			if err == sql.ErrNoRows {
				return []string{}, fmt.Errorf("storage.post-like-dislike.GetPostLikes SQLnoRows:%v", err)
			}
			return nil, fmt.Errorf("storage.post-like-dislike.GetPostLikes Scan:%v", err)
		}
		postLikes = append(postLikes, postLike)
	}
	return postLikes, nil
}

func (d *LikeDislikePostStorage) GetPostLikes(post_id int) ([]string, error) {
	var postLikes []string
	query := `SELECT login FROM likes WHERE post_id = $1`
	rows, err := d.db.Query(query, post_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var postLike string
		if err := rows.Scan(&postLike); err != nil {
			if err == sql.ErrNoRows {
				return []string{}, nil
			}
			return nil, err
		}
		postLikes = append(postLikes, postLike)
	}
	return postLikes, nil
}

func (d *LikeDislikePostStorage) DislikePost(post_id int, login string) error {
	query := `INSERT INTO dislikes(post_id, login) VALUES($1, $2)`
	_, err := d.db.Exec(query, post_id, login)
	if err != nil {
		return fmt.Errorf("storage.DislikePost.DB.Exec : %v", err)
	}
	return nil
}

func (d *LikeDislikePostStorage) LikePost(post_id int, login string) error {
	query := `INSERT INTO likes(post_id, login) VALUES ($1, $2)`
	_, err := d.db.Exec(query, post_id, login)
	if err != nil {
		return fmt.Errorf("storage: like post: %v", err)
	}
	return nil
}

func (d *LikeDislikePostStorage) RemoveDislikeFromPost(post_id int, login string) error {
	query := `DELETE FROM dislikes WHERE post_id = $1 and login = $2`
	_, err := d.db.Exec(query, post_id, login)
	if err != nil {
		return fmt.Errorf("storage: remove disLike from post: %v", err)
	}
	return nil
}

func (d *LikeDislikePostStorage) RemoveLikeFromPost(post_id int, login string) error {
	query := `DELETE FROM likes WHERE post_id = $1 and login = $2`
	_, err := d.db.Exec(query, post_id, login)
	if err != nil {
		return fmt.Errorf("storage: remove like from post: %v", err)
	}
	return nil
}

func (d *LikeDislikePostStorage) PostHasDislike(post_id int, login string) error {
	var query string
	query = `SELECT post_id, login FROM dislikes WHERE post_id = $1 and login = $2`
	row := d.db.QueryRow(query, post_id, login)
	err := row.Scan(&post_id, &login)
	if err != nil {
		return fmt.Errorf("storage.PostHasDislike.QueryRow: %v", err)
	}
	return nil
}

func (d *LikeDislikePostStorage) PostHasLike(post_id int, login string) error {
	var query string
	query = `SELECT post_id, login FROM likes WHERE post_id = $1 and login = $2`
	row := d.db.QueryRow(query, post_id, login)
	err := row.Scan(&post_id, &login)
	if err != nil {
		return fmt.Errorf("storage: post has like: %v", err)
	}
	return nil
}

func (d *LikeDislikePostStorage) AddPostDislike(post_id int, login string) error {
	query := `INSERT INTO dislikes (post_id, login) VALUES ($1, $2)`
	_, err := d.db.Exec(query, post_id, login)
	if err != nil {
		return fmt.Errorf("storage.Like.AddPostDislike - DB.Exec: %v", err)
	}
	return nil
}

func (d *LikeDislikePostStorage) AddPostLike(post_id int, login string) error {
	query := `INSERT INTO likes (post_id, login) VALUES ($1, $2)`
	_, err := d.db.Exec(query, post_id, login)
	if err != nil {
		return fmt.Errorf("storage.Like.AddPostLike - DB.Exec: %v", err)
	}
	return nil
}
