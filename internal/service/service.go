package service

import "forum/internal/storage"

type Service struct {
	Auth
	Post
	Like
	LikeDislikePost
	LikeDislikeComment
}

func NewService(storage *storage.Database) *Service {
	return &Service{
		Auth:               newAuthService(storage.Auth),
		Post:               newPostService(storage.Post),
		Like:               newLikeService(storage.Like),
		LikeDislikePost:    newLikeDislikePostService(storage.LikeDislikePost),
		LikeDislikeComment: newLikeDislikeCommentService(storage.LikeDislikeComment),
	}
}
