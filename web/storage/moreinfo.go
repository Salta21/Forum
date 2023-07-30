package storage

import (
	"database/sql"

	"forum/web/models"
)

type Moreinfo interface {
	GetPostByPostId(postId int) (models.Post, error)
	SelectContentsByPostId(post_id int) ([]string, error)
	CountLikedPosts(postId int) (int, error)
	CountDislikedPosts(postId int) (int, error)
	GetCommentByPostId(post_id int) ([]models.PostTheComment, error)
	GetUserById(n int) (string, string, error)
	CountLikedComment(commentId int) (int, error)
	CountDislikedComment(commentId int) (int, error)
}
type MoreinfoStorage struct {
	db *sql.DB
}

func newMoreinfoStorage(db *sql.DB) *MoreinfoStorage {
	return &MoreinfoStorage{
		db: db,
	}
}

func (s *MoreinfoStorage) GetPostByPostId(postId int) (models.Post, error) {
	rows, err := s.db.Query("SELECT user_id, posts_id,author, title,posts FROM posts WHERE posts_id = ?", postId)
	if err != nil {
		return models.Post{}, err
	}
	var post models.Post
	for rows.Next() {
		err := rows.Scan(&post.User_id, &post.Post_id, &post.Author, &post.Title, &post.Posts)
		if err != nil {
			return models.Post{}, err
		}
		contents, err := s.SelectContentsByPostId(post.Post_id)
		if err != nil {
			return models.Post{}, err
		}
		post.Contents = contents
	}
	return post, nil
}

func (s *MoreinfoStorage) SelectContentsByPostId(post_id int) ([]string, error) {
	var check []string // it is empty
	rows, err := s.db.Query("SELECT content FROM postcontent WHERE post_id = ?", post_id)
	if err != nil {
		return check, err
	}
	var allPost []string
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Content)
		if err != nil {
			return check, err
		}
		allPost = append(allPost, post.Content)
	}
	return allPost, nil
}

func (s *MoreinfoStorage) CountLikedPosts(postId int) (int, error) {
	rows, err := s.db.Query("SELECT COUNT(*) FROM likes WHERE post_id=?", postId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (s *MoreinfoStorage) CountDislikedPosts(postId int) (int, error) {
	rows, err := s.db.Query("SELECT COUNT(*) FROM dislikes WHERE post_id=?", postId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (s *MoreinfoStorage) GetCommentByPostId(post_id int) ([]models.PostTheComment, error) {
	rows, err := s.db.Query("SELECT comment_id,commenter_id,post_id,commenttext FROM comments WHERE post_id = ?", post_id)
	if err != nil {
		return nil, err
	}
	allcomment := []models.PostTheComment{}
	for rows.Next() {
		var comment models.PostTheComment
		err := rows.Scan(&comment.Comment_id, &comment.Commenter_id, &comment.CommentedPost_id, &comment.Commenttext)
		if err != nil {
			return nil, err
		}
		comment.Commentername, _, err = s.GetUserById(comment.Commenter_id)
		if err != nil {
			return nil, err
		}
		allcomment = append(allcomment, comment)
	}
	return allcomment, nil
}

func (s *MoreinfoStorage) GetUserById(n int) (string, string, error) {
	findId := s.db.QueryRow("SELECT * FROM users WHERE  users_id = ?", n) // gets the user_id by token
	user := models.User{}
	err := findId.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return "", "", err
	}
	return user.Username, user.Email, nil
}

func (s *MoreinfoStorage) CountLikedComment(commentId int) (int, error) {
	rows, err := s.db.Query("SELECT COUNT(*) FROM likecomments WHERE comment_id=?", commentId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}

func (s *MoreinfoStorage) CountDislikedComment(commentId int) (int, error) {
	rows, err := s.db.Query("SELECT COUNT(*) FROM dislikecomments WHERE comment_id=?", commentId)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}
