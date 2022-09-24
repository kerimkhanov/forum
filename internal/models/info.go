package models

type Info struct {
	Users             Users
	ProfileUser       Users
	PostsCount        []Posts
	SimilarPosts      []Posts
	Posts             Posts
	Notifications     []Notification
	PostLikes         []string
	PostDislikes      []string
	Comments          []Comments
	CommentsLikes     map[int][]string
	CommentsDislikes  map[int][]string
	CommentsDisCount  int
	CommentsLikeCount int
}
