package service

import (
	"forum/web/models"
	"forum/web/storage"
)

type Moreinfo interface {
	GetPostByPostId(postId int) (models.Post, error)
	CountLikedPosts(postId int) (int, error)
	CountDislikedPosts(postId int) (int, error)
	GetCommentByPostId(post_id int) ([]models.PostTheComment, error)
	CountLikedComment(commentId int) (int, error)
	CountDislikedComment(commentId int) (int, error)
}
type MoreinfoService struct {
	storage storage.Moreinfo
}

func newMoreinfoService(storage storage.Moreinfo) *MoreinfoService {
	return &MoreinfoService{
		storage: storage,
	}
}

func (s *MoreinfoService) GetPostByPostId(postId int) (models.Post, error) {
	posts, err := s.storage.GetPostByPostId(postId)
	if err != nil {
		return models.Post{}, err
	}
	return posts, nil
}

func (s *MoreinfoService) CountLikedPosts(postId int) (int, error) {
	amount, err := s.storage.CountLikedPosts(postId)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func (s *MoreinfoService) CountDislikedPosts(postId int) (int, error) {
	amount, err := s.storage.CountDislikedPosts(postId)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func (s *MoreinfoService) GetCommentByPostId(post_id int) ([]models.PostTheComment, error) {
	comments, err := s.storage.GetCommentByPostId(post_id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *MoreinfoService) CountLikedComment(commentId int) (int, error) {
	amount, err := s.storage.CountLikedComment(commentId)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func (s *MoreinfoService) CountDislikedComment(commentId int) (int, error) {
	amount, err := s.storage.CountDislikedComment(commentId)
	if err != nil {
		return 0, err
	}
	return amount, nil
}
