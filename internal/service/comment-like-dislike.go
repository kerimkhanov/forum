package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"forum/internal/storage"
)

type LikeDislikeComment interface {
	AddCommentLike(comment_id int, login string) error
	GetPostIdByComment(comment_id int, login string) (int, error)
	AddCommentDislike(comment_id int, login string) error
	GetCommentLikes(post_id int, login string) (map[int][]string, error)
	GetCommnetDislikes(post_id int, login string) (map[int][]string, error)
}

type LikeDislikeCommentService struct {
	storage storage.LikeDislikeComment
}

func NewServiceLikeDislikeComment(s storage.LikeDislikeComment) *Service {
	log.Println("NewServiceLike implementation")
	return &Service{
		LikeDislikeComment: newLikeDislikeCommentService(s),
	}
}

func newLikeDislikeCommentService(storage storage.LikeDislikeComment) *LikeDislikeCommentService {
	log.Println("newPostService implementation")
	return &LikeDislikeCommentService{
		storage: storage,
	}
}

func (s *LikeDislikeCommentService) GetCommnetDislikes(post_id int, login string) (map[int][]string, error) {
	Users, err := s.storage.GetCommentDislikes(post_id, login)
	if err != nil {
		return nil, fmt.Errorf("storage.likeDislikeComments.GetCommentDislikes: %v", err)
	}
	return Users, nil
}

func (s *LikeDislikeCommentService) GetCommentLikes(post_id int, login string) (map[int][]string, error) {
	Users, err := s.storage.GetCommentLikes(post_id, login)
	if err != nil {
		return nil, fmt.Errorf("service.comment-like-dislike.GetCommentLikes:%w", err)
	}
	return Users, nil
}

func (s *LikeDislikeCommentService) GetPostIdByComment(comment_id int, login string) (int, error) {
	return s.storage.GetPostIdByComment(comment_id, login)
}

func (s *LikeDislikeCommentService) AddCommentDislike(comment_id int, login string) error {
	if err := s.storage.CommentHasDislike(comment_id, login); err == nil {
		if err := s.storage.RemoveDislikeFromComment(comment_id, login); err != nil {
			fmt.Println("CommentHasDislike - removeDislikeFromCommnet")
			return fmt.Errorf("service: dislike comment: %v", err)
		}
		fmt.Println("CommentHasDislike - removeDislikeFromCommnet = nil")
		return nil
	} else if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("add commentdislike - sql.ErrNoRows")
		return fmt.Errorf("service: dislike comment: %v", err)
	}
	if err := s.storage.CommentHasLike(comment_id, login); err == nil {
		if err := s.storage.RemoveLikeFromComment(comment_id, login); err != nil {
			fmt.Println("CommentHaslike - removelikeFromCommnet")
			return fmt.Errorf("service: dislike comment: %v", err)
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("add commentdislike - sql.ErrNoRows2")
		return fmt.Errorf("service: dislike comment: %w", err)
	}
	if err := s.storage.AddCommentDislike(comment_id, login); err != nil {
		fmt.Println("AddCommentDislike - real")
		return fmt.Errorf("service: dislike comment: %w", err)
	}
	return nil
}

func (s *LikeDislikeCommentService) AddCommentLike(comment_id int, login string) error {
	if err := s.storage.CommentHasLike(comment_id, login); err == nil {
		if err := s.storage.RemoveLikeFromComment(comment_id, login); err != nil {
			fmt.Println("Check Error HasDislikeComment or RemoveDislikeComment")
			return fmt.Errorf("check error CommentLike: %v", err)
		}
		fmt.Println("Check Error CommentDislike or RemoveDislike")
		return nil
	} else if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("service: comment Dislike post: %v", err)
		return fmt.Errorf("service: dislike Comment: %v", err)
	}
	if err := s.storage.CommentHasDislike(comment_id, login); err == nil {
		if err := s.storage.RemoveDislikeFromComment(comment_id, login); err != nil {
			return fmt.Errorf("Check error CommentDislike: %v", err)
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Check Error HasCommentDislike or RemoveCommentDislike")
		return fmt.Errorf("service: dislike comment %v", err)
	}
	if err := s.storage.LikeComment(comment_id, login); err != nil {
		fmt.Println("check Like Comment -- good WORKING")
		return fmt.Errorf("service.AddComment Like :%v", err)
	}
	fmt.Println("Here1")
	return nil
}
