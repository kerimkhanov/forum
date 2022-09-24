package service

import (
	"fmt"
	"log"

	"forum/internal/models"
	"forum/internal/storage"
)

type Post interface {
	CreatePost(title, body, author string, tags []string) error
	GetAllPosts() ([]models.Posts, error)
	GetPostById(id int) (models.Posts, error)
	AddComment(id int, text, login string) error
}

type PostService struct {
	storage storage.Post
}

func NewServicePost(s storage.Post) *Service {
	log.Println("NewServicePost implementation")
	return &Service{
		Post: newPostService(s),
	}
}

func newPostService(storage storage.Post) *PostService {
	log.Println("newPostService implementation")
	return &PostService{
		storage: storage,
	}
}

func (s *PostService) AddComment(id int, text, login string) error {
	return s.storage.AddComment(id, text, login)
}

func (s *PostService) CreatePost(title, body, author string, tags []string) error {
	fmt.Println("------------------------------------------")
	fmt.Sprintf("\n%s - title, %s - body, %s - author, %s - tags --- > \n", title, body, author, tags)
	var post models.Posts = models.Posts{
		Title:  title,
		Text:   body,
		Author: author,
		Tags:   tags,
	}
	return s.storage.CreatePost(post)
}

func (s *PostService) GetAllPosts() ([]models.Posts, error) {
	return s.storage.GetAllPosts()
}

func (s *PostService) GetPostById(id int) (models.Posts, error) {
	return s.storage.GetPostById(id)
}
