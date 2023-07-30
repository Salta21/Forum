package storage

import (
	"database/sql"

	"forum/web/models"
)

type Comment interface {
	FillTheCommentTable(commenter_id int, post_id int, commenttext string) error
	IfLikerLikedComment(liker_id int, commentId int) (bool, error)
	DeleteLikeComments(liker_id, commentId int) error
	IfDislikerDislikedComment(disliker_id int, commentId int) (bool, error)
	DeleteDislikeComments(disliker_id, commentId int) error
	FillTheDislikeCommentsTable(disliker_id, commentId int) error
	FillTheLikeCommentsTable(liker_id, comment_id int) error
	GetPostbyPost_id(post_id int) error
}
type CommentStorage struct {
	db *sql.DB
}

func newCommentStorage(db *sql.DB) *CommentStorage {
	return &CommentStorage{
		db: db,
	}
}

func (s *CommentStorage) FillTheCommentTable(commenter_id int, post_id int, commenttext string) error {
	records := `INSERT INTO comments(commenter_id, post_id,commenttext) VALUES (?,?,?)` // the query itself
	query, err := s.db.Prepare(records)                                                 // prepares above-mentioned query INSERT INTO users(token) VALUES(?) WHERE users_id = ?
	if err != nil {
		return err
	}
	_, err = query.Exec(commenter_id, post_id, commenttext)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentStorage) IfLikerLikedComment(liker_id int, commentId int) (bool, error) { // checkforduplicate
	find_like := s.db.QueryRow("SELECT liker_id, comment_id FROM likecomments WHERE liker_id = ? AND comment_id = ?", liker_id, commentId) // gets the user_id by token
	comment := models.Comment{}
	err := find_like.Scan(&comment.Liker_id, &comment.Comment_id)
	if err != nil {
		return false, nil // the liker didnt like the comment earlier
	}
	return true, nil
}

func (s *CommentStorage) DeleteLikeComments(liker_id, commentId int) error {
	_, err := s.db.Exec("DELETE FROM likecomments WHERE liker_id=? AND comment_id=?", liker_id, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentStorage) IfDislikerDislikedComment(disliker_id int, commentId int) (bool, error) { // checkforduplicate
	find_like := s.db.QueryRow("SELECT disliker_id, comment_id FROM dislikecomments WHERE disliker_id = ? AND comment_id = ?", disliker_id, commentId) // gets the user_id by token
	user := models.Comment{}
	err := find_like.Scan(&user.Disliker_id, &user.Comment_id)
	if err != nil {
		return false, nil // the liker didnt like the comment earlier
	}
	return true, nil
}

func (s *CommentStorage) DeleteDislikeComments(disliker_id, commentId int) error {
	_, err := s.db.Exec("DELETE FROM dislikecomments WHERE disliker_id=? AND comment_id=?", disliker_id, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentStorage) FillTheDislikeCommentsTable(disliker_id, commentId int) error {
	records := `INSERT INTO dislikecomments(disliker_id, comment_id) VALUES (?,?)`
	query, err := s.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(disliker_id, commentId) // executes this query
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentStorage) FillTheLikeCommentsTable(liker_id, comment_id int) error {
	records := `INSERT INTO likecomments(liker_id, comment_id) VALUES (?,?)`
	query, err := s.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(liker_id, comment_id) // executes this query
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentStorage) GetPostbyPost_id(post_id int) error { // you need to finish this function by getting all the posts by this id and display them
	find_like := s.db.QueryRow("SELECT posts_id FROM posts WHERE posts_id = ?", post_id) // gets the user_id by token
	comment := models.Post{}
	err := find_like.Scan(&comment.Post_id)
	if err != nil {
		return err
	}
	return nil
}
