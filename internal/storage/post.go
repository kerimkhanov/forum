package storage

import (
	"database/sql"
	"fmt"

	"forum/internal/models"
)

type Post interface {
	CreatePost(post models.Posts) error
	GetAllPosts() ([]models.Posts, error)
	GetPostById(id int) (models.Posts, error)
	AddComment(id int, text, login string) error
	GetPostsByCategory(category string) ([]models.Posts, error)
}

type PostStorage struct {
	db *sql.DB
}

func newPostStorage(db *sql.DB) *PostStorage {
	return &PostStorage{
		db: db,
	}
}

func (d *PostStorage) GetPostsByCategory(category string) ([]models.Posts, error) {
	var posts []models.Posts
	query := fmt.Sprintf("SELECT * FROM posts WHERE id IN (SELECT post_id FROM tags WHERE tags = $1)", category)
	rows, err := d.db.Query(query, &category)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post models.Posts
		if err := rows.Scan(&post.Id, &post.Title, &post.Text, &post.Author, &post.CreatedAt, &post.Likes, &post.Dislikes, &post.Comments); err != nil {
			return nil, fmt.Errorf("storage.post.GetPostByCategory", err)
		}
		posts = append(posts, post)

	}
	return posts, nil
}

func (d *PostStorage) AddComment(id int, text, login string) error {
	fmt.Println("Here1")
	query := "INSERT INTO comments (post_id, comment, login) VALUES ($1, $2, $3)"
	_, err := d.db.Exec(query, id, text, login)
	if err != nil {
		fmt.Println("Here2")
		return fmt.Errorf("storage.AddComment - DB.Exec: %v", err)
	}
	return nil
}

func (d *PostStorage) GetPostById(id int) (models.Posts, error) {
	onePost := models.Posts{}
	query := "SELECT id, title, text, author FROM post WHERE id = $1"
	row := d.db.QueryRow(query, id)
	err := row.Scan(&onePost.Id, &onePost.Title, &onePost.Text, &onePost.Author)
	if err != nil {
		return models.Posts{}, fmt.Errorf("storage.GetPostById - row.Scan: %v", err)
	}
	tagQuery := "SELECT tag FROM tags WHERE post_id = $1"
	tagRow, err := d.db.Query(tagQuery, id)
	if err != nil {
		return models.Posts{}, fmt.Errorf("storage.GetPostById - db.query - tagRow: %v", err)
	}
	for tagRow.Next() {
		var tag string
		err = tagRow.Scan(&tag)
		if err != nil {
			return models.Posts{}, fmt.Errorf("storage.GetPostById - row.Scan - tagRow: %v", err)
		}
		onePost.Tags = append(onePost.Tags, tag)
	}
	commQuery := "SELECT comment_id, post_id, comment, login FROM comments WHERE post_id = $1"
	comRow, err := d.db.Query(commQuery, id)
	if err != nil {
		return models.Posts{}, fmt.Errorf("storage.GetPostById - DB.Query - comRow: %v", err)
	}
	for comRow.Next() {
		var com models.Comments
		err = comRow.Scan(&com.Comment_id, &com.Post_id, &com.Comment, &com.Login)
		if err != nil {
			return models.Posts{}, fmt.Errorf("storage.GetPostById - row.Scan - tagRow: %v", err)
		}
		onePost.Comments = append(onePost.Comments, com)
	}
	return onePost, nil
}

func (d *PostStorage) GetAllPosts() ([]models.Posts, error) {
	post := []models.Posts{}
	query := "SELECT id, title, text, author FROM post"
	row, err := d.db.Query(query)
	if err != nil {
		return []models.Posts{}, fmt.Errorf("storage.GetAllPosts - DB.Query: %v", err)
	}
	for row.Next() {
		onePost := models.Posts{}
		err = row.Scan(&onePost.Id, &onePost.Title, &onePost.Text, &onePost.Author)
		if err != nil {
			return []models.Posts{}, fmt.Errorf("storage.GetAllPosts - row.Scan %v", err)
		}
		tagQuery := "SELECT tag FROM tags WHERE post_Id = $1"
		tagRow, err := d.db.Query(tagQuery, onePost.Id)
		if err != nil {
			return []models.Posts{}, fmt.Errorf("storage.GetAllPosts - db.Query - tagRow: %v", err)
		}
		for tagRow.Next() {
			var tag string
			err = tagRow.Scan(&tag)
			if err != nil {
				return []models.Posts{}, fmt.Errorf("storage.GetAllPosts - row.Scan - tagRow: %v", err)
			}
			onePost.Tags = append(onePost.Tags, tag)
		}
		post = append(post, onePost)
	}
	return post, nil
}

func (d *PostStorage) CreatePost(post models.Posts) error {
	query := "INSERT INTO Post (title, text, author) values(?, ?, ?)"
	fmt.Printf("\n Create post -- author: %s \n", post.Author)
	sqlResult, err := d.db.Exec(query, post.Title, post.Text, post.Author)
	if err != nil {
		return fmt.Errorf("storage.post.go - CreatePost")
	}
	query = "INSERT INTO Tags (post_id, tag) values($1, $2)"
	postId, err := sqlResult.LastInsertId()
	if err != nil {
		return fmt.Errorf("storage.post.go create post sqlResult")
	}
	for _, w := range post.Tags {
		fmt.Println(w)
		fmt.Println(postId)
		_, err := d.db.Exec(query, postId, w)
		if err != nil {
			return fmt.Errorf("storage.post.go - CreatePost: %v", err)
		}
	}
	return nil
}

func (d *PostStorage) DeletePost(post models.Posts) (error, models.Posts) {
	query := "DELETE FROM Posts WHERE "
	_, err := d.db.Exec(query, post.Title, post.Text, post.Author)
	if err != nil {
		return fmt.Errorf("storage.post.go - CreatePost"), models.Posts{}
	}
	return fmt.Errorf(""), post
}
