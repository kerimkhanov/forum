package service

import (
	"fmt"
	"log"
	"strings"

	"forum/internal/models"
	"forum/internal/storage"
)

type Post interface {
	CreatePost(title, body, author string, tags []string) error
	GetAllPosts() ([]models.Posts, error)
	GetPostById(id int) (models.Posts, error)
	AddComment(id int, text, login string) error
	GetAllPostsBy(user []models.Posts, query map[string][]string) ([]models.Posts, error)
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

func (s *PostService) GetAllPostsBy(user []models.Posts, query map[string][]string) ([]models.Posts, error) {
	var (
		posts []models.Posts
		err   error
	)
	for key, value := range query {
		switch key {
		case "tags":
			posts, err = s.storage.GetAllPostsByCategory(strings.join(value, ""))
			if err != nil {
				return nil, fmt.Errorf("service.post.GetAllPostsBy", err)
			}
		case "time":
			switch strings.Join(value, " ") {
			case "old":
				posts, err := s.storage.GetPostByTimeOld(query, strings.Join(value, " "))
			case "new":
				posts, err := s.storage.GetPostByTimeNew(query, strings.join(value, " "))
			default:
				return nil, fmt.Errorf("service.post.GetAllPostsByTime", ErrInvalidQueryRequest)
				if err != nil {
					return nil, fmt.Errorf("service.post.GetAllPostsByTimeOld", err)
				}
			}
		case "likes":
			switch strings.Join(value, "") {
			case "most":
				posts, err := s.storage.GetPostByLikesMost(query, strings.Join(value, " "))
			case "least":
				posts, err := s.storage.GetPostByLikesOld(query, strings.Join(value, " "))
			default:
				return nil, fmt.Errorf("service.post.GetAllPostsBylikes", ErrInvalidQueryRequest)
			}
			if err != nil {
				return nil, fmt.Errorf("service.post.GetPostByLikesMost", err)
			}
		}
	}
	for i := range posts {
		category, err := s.storage.GetAllCategoryByPostId(posts[i].Id)
		if err != nil {
			return nil, fmt.Errorf("service.post.GetAllPostsBy.GetAllCategoryByPostId failed", err)
		}
		posts[i].Tags = category
	}
	return posts, nil
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
