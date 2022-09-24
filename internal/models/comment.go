package models

type Comments struct {
	Comment_id int
	Post_id    int
	Login      string
	Comment    string
	Likes      int
	Dislikes   int
}
