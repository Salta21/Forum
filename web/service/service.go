package service

import "forum/web/storage"

type Service struct {
	Auth
	Profile
	Moreinfo
	Comment
	User
	Post
}

func NewService(storages *storage.Storage) *Service {
	return &Service{
		Auth:     newAuthService(storages.Auth),
		Profile:  newProfileService(storages.Profile),
		Moreinfo: newMoreinfoService(storages.Moreinfo),
		Comment:  newCommentService(storages.Comment),
		User:     newUserService(storages.User),
		Post:     newPostService(storages.Post),
	}
}
