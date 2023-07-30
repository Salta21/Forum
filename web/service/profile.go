package service

import (
	"forum/web/models"
	"forum/web/storage"
)

type Profile interface {
	GetUserService(userID int) (string, string, error)
	GetAllPostStorage() ([]models.Post, error)
	GetUserById(n int) (string, string, error)
	FillThePostTable(userId int, post string, username string, title string) error
	FindPostIdbyPost(posts string) (int, error)
	FillTheContentTable(post_id int, contents []string) error
	GetAllPostsByCategory(c string) ([]models.Post, error)
}

type ProfileService struct {
	storage storage.Profile
}

func newProfileService(storage storage.Profile) *ProfileService {
	return &ProfileService{
		storage: storage,
	}
}

func (s *ProfileService) GetUserService(userID int) (string, string, error) {
	username, email, err := s.storage.GetUserById(userID)
	if err != nil {
		return "", "", err
	}
	return username, email, nil
}

func (s *ProfileService) GetAllPostStorage() ([]models.Post, error) {
	allPost, err := s.storage.GetAllPost()
	if err != nil {
		return nil, err
	}
	return allPost, nil
}

func (s *ProfileService) GetUserById(n int) (string, string, error) {
	username, email, err := s.storage.GetUserById(n)
	if err != nil {
		return "", "", err
	}
	return username, email, nil
}

func (s *ProfileService) FillThePostTable(userId int, post string, username string, title string) error {
	if err := s.storage.FillThePostTable(userId, post, username, title); err != nil {
		return err
	}
	return nil
}

func (s *ProfileService) FindPostIdbyPost(posts string) (int, error) {
	post_id, err := s.storage.FindPostIdbyPost(posts)
	if err != nil {
		return 0, err
	}
	return post_id, nil
}

func (s *ProfileService) FillTheContentTable(post_id int, contents []string) error {
	for d := 0; d < len(contents); d++ {
		if err := s.storage.FillTheContentTable(post_id, contents[d]); err != nil {
			return err
		}
	}
	return nil
}

func (s *ProfileService) GetAllPostsByCategory(c string) ([]models.Post, error) {
	allPost, err := s.storage.GetAllPostsByCategory(c)
	if err != nil {
		return nil, err
	}
	return allPost, nil
}
