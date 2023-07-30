package service

import "forum/web/storage"

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

type CommentService struct {
	storage storage.Comment
}

func newCommentService(storage storage.Comment) *CommentService {
	return &CommentService{
		storage: storage,
	}
}

func (s *CommentService) FillTheCommentTable(commenter_id int, post_id int, commenttext string) error {
	if err := s.storage.FillTheCommentTable(commenter_id, post_id, commenttext); err != nil {
		return err
	}
	return nil
}

func (s *CommentService) IfLikerLikedComment(liker_id int, commentId int) (bool, error) {
	check, err := s.storage.IfLikerLikedComment(liker_id, commentId)
	return check, err
}

func (s *CommentService) DeleteLikeComments(liker_id, commentId int) error {
	if err := s.storage.DeleteLikeComments(liker_id, commentId); err != nil {
		return err
	}
	return nil
}

func (s *CommentService) IfDislikerDislikedComment(disliker_id int, commentId int) (bool, error) {
	check, err := s.storage.IfDislikerDislikedComment(disliker_id, commentId)
	return check, err
}

func (s *CommentService) DeleteDislikeComments(disliker_id, commentId int) error {
	if err := s.storage.DeleteDislikeComments(disliker_id, commentId); err != nil {
		return err
	}
	return nil
}

func (s *CommentService) FillTheDislikeCommentsTable(disliker_id, commentId int) error {
	if err := s.storage.FillTheDislikeCommentsTable(disliker_id, commentId); err != nil {
		return err
	}
	return nil
}

func (s *CommentService) FillTheLikeCommentsTable(liker_id, comment_id int) error {
	if err := s.storage.FillTheLikeCommentsTable(liker_id, comment_id); err != nil {
		return err
	}
	return nil
}

func (s *CommentService) GetPostbyPost_id(post_id int) error { // you need to finish this function by getting all the posts by this id and display them
	err := s.storage.GetPostbyPost_id(post_id)
	return err
}
