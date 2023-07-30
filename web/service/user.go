package service

import "forum/web/storage"

type User interface {
	GetUserIdByToken(userToken string) (int, error)
}
type UserService struct {
	storage storage.User
}

func newUserService(storage storage.User) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s *UserService) GetUserIdByToken(userToken string) (int, error) {
	user_id, err := s.storage.GetUserIdByToken(userToken)
	if err != nil {
		return 0, err
	}
	return user_id, nil
}
