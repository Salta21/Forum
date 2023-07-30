package storage

import (
	"database/sql"

	"forum/web/models"
)

type User interface {
	GetUserIdByToken(userToken string) (int, error)
}
type UserStorage struct {
	db *sql.DB
}

func newUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) GetUserIdByToken(userToken string) (int, error) {
	findId := s.db.QueryRow("SELECT user_id FROM sessions WHERE  token = ?", userToken) // gets the user_id by token
	session := models.Session{}
	err := findId.Scan(&session.User_id)
	if err != nil {
		return 0, err
	}
	return session.User_id, nil
}
