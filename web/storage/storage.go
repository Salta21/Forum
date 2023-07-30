package storage

import "database/sql"

type Storage struct {
	Auth
	Profile
	Moreinfo
	Comment
	User
	Post
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Auth:     newAuthStorage(db),
		Profile:  newProfileStorage(db),
		Moreinfo: newMoreinfoStorage(db),
		Comment:  newCommentStorage(db),
		User:     newUserStorage(db),
		Post:     newPostStorage(db),
	}
}
