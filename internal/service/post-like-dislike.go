package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"forum/internal/storage"
)

type LikeDislikePost interface {
	LikePost(post_id int, login string) error
	DislikePost(post_id int, login string) error
	GetPostLikes(post_id int) ([]string, error)
	GetPostDislikes(post_id int) ([]string, error)
}

type LikeDislikePostService struct {
	storage storage.LikeDislikePost
}

func NewServiceLikeDislikePost(s storage.LikeDislikePost) *Service {
	log.Println("NewServiceLike implementation")
	return &Service{
		LikeDislikePost: newLikeDislikePostService(s),
	}
}

func newLikeDislikePostService(storage storage.LikeDislikePost) *LikeDislikePostService {
	log.Println("newPostService implementation")
	return &LikeDislikePostService{
		storage: storage,
	}
}

func (s *LikeDislikePostService) GetPostDislikes(post_id int) ([]string, error) {
	Users, err := s.storage.GetPostDislikes(post_id)
	if err != nil {
		return nil, fmt.Errorf("service: get post likes: %v")
	}
	return Users, nil
}

func (s *LikeDislikePostService) GetPostLikes(post_id int) ([]string, error) {
	Users, err := s.storage.GetPostLikes(post_id)
	if err != nil {
		return nil, fmt.Errorf("service.GetPostLikes: %v", err)
	}
	return Users, nil
}

func (s *LikeDislikePostService) DislikePost(post_id int, login string) error {
	if err := s.storage.PostHasDislike(post_id, login); err == nil {
		if err := s.storage.RemoveDislikeFromPost(post_id, login); err != nil {
			return fmt.Errorf("service: dislike post: %v", err)
		}
		fmt.Println("Check Error HasDisLike or RemoveDislLikeFromPost")
		return nil
	} else if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Check Error HasLike or RemoveLikeFromPost WTF ?")
		return fmt.Errorf("service: dislike post %v", err)
	}
	if err := s.storage.PostHasLike(post_id, login); err == nil {
		if err := s.storage.RemoveLikeFromPost(post_id, login); err != nil {
			fmt.Println("Check Error HasDisLike or RemoveDisLikeFromPost")
			return fmt.Errorf("service: dislike post %v", err)
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Check Error HasLike or RemoveLikeFromPost WTF2?")

		return fmt.Errorf("service: dislike post: %v", err)
	}
	if err := s.storage.DislikePost(post_id, login); err != nil {
		fmt.Println("Check Error LikePost--GOOD WORKING")

		return fmt.Errorf("service: dislike post: %v", err)
	}
	fmt.Println("Here1")
	return nil
}

func (s *LikeDislikePostService) LikePost(post_id int, login string) error {
	fmt.Println("sam petuh2LikePost.STorage")
	if err := s.storage.PostHasLike(post_id, login); err == nil {
		if err := s.storage.RemoveLikeFromPost(post_id, login); err != nil {
			return fmt.Errorf("service: Like post: %v", err)
		}
		fmt.Println("Check Error HasLike or RemoveLikeFromPost")
		return nil
	} else if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Check Error HasLike or RemoveLikeFromPost WTF ?")
		return fmt.Errorf("service: like post: %v", err)
	}
	if err := s.storage.PostHasDislike(post_id, login); err == nil {
		if err := s.storage.RemoveDislikeFromPost(post_id, login); err != nil {
			fmt.Println("Check Error HasDisLike or RemoveDisLikeFromPost")
			return fmt.Errorf("service: Like post: %v", err)
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Check Error HasLike or RemoveLikeFromPost WTF2?")
		return fmt.Errorf("service: like post: %v", err)
	}
	if err := s.storage.LikePost(post_id, login); err != nil {
		fmt.Println("Check Error LikePost--GOOD WORKING")
		return fmt.Errorf("service: like post: %v", err)
	}
	fmt.Println("Here2")
	return nil
}
