package service

import (
	"log"

	"forum/internal/storage"
)

type Like interface{}

type LikeService struct {
	storage storage.Like
}

func NewServiceLike(s storage.Like) *Service {
	log.Println("NewServiceLike implementation")
	return &Service{
		Like: newLikeService(s),
	}
}

func newLikeService(storage storage.Like) *LikeService {
	log.Println("newPostService implementation")
	return &LikeService{
		storage: storage,
	}
}
