package models

type PostTheComment struct {
	Post_id          int
	Comment_id       int
	Commenter_id     int
	CommentedPost_id int
	Commenttext      string
	Commentername    string
	Likes            int
	Dislikes         int
}
