package storage

import (
	"database/sql"
	"fmt"
)

type LikeDislikeComment interface {
	CommentHasLike(comment_id int, login string) error
	RemoveLikeFromComment(comment_id int, login string) error
	CommentHasDislike(comment_id int, login string) error
	RemoveDislikeFromComment(comment_id int, login string) error
	LikeComment(comment_id int, login string) error
	GetPostIdByComment(comment_id int, login string) (int, error)
	AddCommentDislike(comment_id int, login string) error
	GetCommentLikes(post_id int, login string) (map[int][]string, error)
	GetCommentDislikes(post_id int, login string) (map[int][]string, error)

	// GetCommnetDislikes(comment_id int, login string) (map[int][]string, error)
}

type LikeDislikeCommentStorage struct {
	db *sql.DB
}

func newLikeDislikeCommentStorage(db *sql.DB) *LikeDislikeCommentStorage {
	return &LikeDislikeCommentStorage{
		db: db,
	}
}

// func (d *LikeDislikeCommentStorage) GetCommnetDislikes(comment_id int, login string) (map[int][]string, error) {
// 	queryForCommentsId := `SELECT comment_id FROM dislikes WHERE login = $1`
// }

func (d *LikeDislikeCommentStorage) GetCommentDislikes(post_id int, login string) (map[int][]string, error) {
	queryForCommentsId := `SELECT comment_id FROM comments WHERE post_id = $1`
	queryForUsers := `SELECT login FROM dislikes WHERE comment_id = $1`
	users := make(map[int][]string)
	rowsComment, err := d.db.Query(queryForCommentsId, post_id)
	if err != nil {
		return nil, fmt.Errorf("storage.comment-like-dislike.GetCommentDislikes.Query: %v", err)
	}
	for rowsComment.Next() {
		var id int
		if err := rowsComment.Scan(&id); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, fmt.Errorf("storage.comment-like-dislike.GetCommentDislikes2.Query: %v", err)
		}
		var usernames []string
		rowsUsers, err := d.db.Query(queryForUsers, id)
		if err != nil {
			return nil, fmt.Errorf("storage.comment-like-dislike.GetCommentLikes.QueryForUsers.Query: %v", err)
		}
		for rowsUsers.Next() {
			var username string
			if err := rowsUsers.Scan(&username); err != nil {
				return nil, fmt.Errorf("storage.comment-like-dislike.GetCommentDislikes.Scan :%v", err)
			}
			usernames = append(usernames, username)
		}
		users[id] = usernames
	}
	return users, nil
}

func (d *LikeDislikeCommentStorage) GetCommentLikes(post_id int, login string) (map[int][]string, error) {
	queryForCommentsId := `SELECT comment_id FROM comments WHERE post_id = $1`
	queryForUseres := `SELECT login FROM likes WHERE comment_id = $1`
	users := make(map[int][]string)
	rowsComment, err := d.db.Query(queryForCommentsId, post_id)
	if err != nil {
		return nil, fmt.Errorf("storage.LikeDislikeComments.GetCommentLikes: %v", err)
	}
	for rowsComment.Next() {
		var id int
		err = rowsComment.Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, fmt.Errorf("storage.GetCommentLikes: %v", err)
		}
		var usernames []string
		rowsUsers, err := d.db.Query(queryForUseres, id)
		if err != nil {
			return nil, fmt.Errorf("storage.comment-like-dislike.GetCommentLikes.QueryForUsers: %v", err)
		}
		for rowsUsers.Next() {
			var username string
			if err := rowsUsers.Scan(&username); err != nil {
				return nil, fmt.Errorf("storage.comment-like-dislike.GetCommentLikes.QueryForUsers: %v", err)
			}
			usernames = append(usernames, username)
		}
		users[id] = usernames
	}
	return users, nil
}

func (d *LikeDislikeCommentStorage) GetPostIdByComment(comment_id int, login string) (int, error) {
	fmt.Printf("comment_id: % , login: %s", comment_id, login)
	query := `SELECT post_id FROM comments WHERE comment_id = $1`
	var post_id int
	row := d.db.QueryRow(query, comment_id)
	err := row.Scan(&post_id)
	if err != nil {
		return 0, fmt.Errorf("storage.GetPostIdByComment.QueryRow: %v", err)
	}
	return post_id, nil
}

func (d *LikeDislikeCommentStorage) AddCommentDislike(comment_id int, login string) error {
	query := `INSERT INTO dislikes(comment_id, login) VALUES ($1, $2)`
	_, err := d.db.Exec(query, comment_id, login)
	if err != nil {
		return fmt.Errorf("storage.Like.AddPostLike - DB.Exec: %v", err)
	}
	return nil
}

func (d *LikeDislikeCommentStorage) LikeComment(comment_id int, login string) error {
	query := `INSERT INTO likes(comment_id, login) VALUES ($1, $2)`
	_, err := d.db.Exec(query, comment_id, login)
	if err != nil {
		return fmt.Errorf("storage: like comment: %v", err)
	}
	return nil
}

func (d *LikeDislikeCommentStorage) RemoveDislikeFromComment(comment_id int, login string) error {
	query := `DELETE FROM dislikes WHERE comment_id = $1 and login = $2`
	_, err := d.db.Exec(query, comment_id, login)
	if err != nil {
		return fmt.Errorf("storage: remove dislike from comment: %v", err)
	}
	return nil
}

func (d *LikeDislikeCommentStorage) RemoveLikeFromComment(comment_id int, login string) error {
	query := `DELETE FROM likes WHERE comment_id = $1 and login = $2`
	_, err := d.db.Exec(query, comment_id, login)
	if err != nil {
		return fmt.Errorf("storage: remove like from comment: %v", err)
	}
	return nil
}

func (d *LikeDislikeCommentStorage) CommentHasDislike(comment_id int, login string) error {
	var query string
	query = `SELECT comment_id, login FROM dislikes WHERE comment_id = $1 and login = $2`
	row := d.db.QueryRow(query, comment_id, login)
	err := row.Scan(&comment_id, &login)
	if err != nil {
		return fmt.Errorf("storage: postComment like: %v", err)
	}
	return nil
}

func (d *LikeDislikeCommentStorage) CommentHasLike(comment_id int, login string) error {
	var query string
	query = `SELECT comment_id, login FROM likes WHERE comment_id = $1 and login = $2`
	row := d.db.QueryRow(query, comment_id, login)
	err := row.Scan(&comment_id, &login)
	if err != nil {
		return fmt.Errorf("storage: postComment Like: %v", err)
	}
	return nil
}
