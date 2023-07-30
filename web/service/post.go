package service

import (
	"forum/web/models"
	_ "forum/web/models"
	"forum/web/storage"
)

type Post interface {
	GetPostsById(user_id int) ([]models.Post, error)
	GetLikedPostsById(user_id int) ([]models.LikedPosts, error)
	IfUserLikedPost(user_id string, post_id string) (bool, error)
	DeleteLikedPost(usId, postId int) error
	IfUserDislikedPost(user_id string, post_id string) (bool, error)
	DeleteDislikedPost(usId, postId int) error
	FillTheDislikesPostTable(usId, postId int) error
	FillTheLikesPostTable(usId, postId int) error
}

type PostService struct {
	storage storage.Post
}

func newPostService(storage storage.Post) *PostService {
	return &PostService{
		storage: storage,
	}
}

func (s *PostService) GetPostsById(user_id int) ([]models.Post, error) { // you need to finish this function by getting all the posts by this id and display them
	posts, err := s.storage.GetPostsById(user_id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetLikedPostsById(user_id int) ([]models.LikedPosts, error) {
	likedPosts, err := s.storage.GetLikedPostsById(user_id)
	if err != nil {
		return nil, err
	}
	return likedPosts, nil
}

func (s *PostService) IfUserLikedPost(user_id string, post_id string) (bool, error) {
	check, err := s.storage.IfUserLikedPost(user_id, post_id)
	return check, err
}

func (s *PostService) DeleteLikedPost(usId, postId int) error {
	if err := s.storage.DeleteLikedPost(usId, postId); err != nil {
		return err
	}
	return nil
}

func (s *PostService) IfUserDislikedPost(user_id string, post_id string) (bool, error) {
	check, err := s.storage.IfUserDislikedPost(user_id, post_id)
	return check, err
}

func (s *PostService) DeleteDislikedPost(usId, postId int) error {
	if err := s.storage.DeleteDislikedPost(usId, postId); err != nil {
		return err
	}
	return nil
}

func (s *PostService) FillTheDislikesPostTable(usId, postId int) error {
	if err := s.storage.FillTheDislikesPostTable(usId, postId); err != nil {
		return err
	}
	return nil
}

func (s *PostService) FillTheLikesPostTable(usId, postId int) error {
	if err := s.storage.FillTheLikesPostTable(usId, postId); err != nil {
		return err
	}
	return nil
}
