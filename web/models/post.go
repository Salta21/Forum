package models

type Post struct {
	Comment_id    int
	Commenter_id  int
	Author        string
	Username      string
	Post_id       int
	User_id       int
	Likes         int
	Dislikes      int
	Posts         string
	Title         string
	Content       string
	Contents      []string
	Commenttext   string
	Commentername string
}
