package storage

import (
	"database/sql"
	"fmt"
	"time"

	"forum/web/models"
)

type Auth interface {
	CreateUser(user models.User) error
	GetUser(user models.User) (models.User, error)
	CheckEmailforDuplicate(user models.User) error
	CheckIfTheSessionExists(ID int64) (bool, error)
	DeleteSession(ID int64) error
	SaveSession(ID int64, token string, cookieTime time.Time) error
}

type AuthStorage struct {
	db *sql.DB
}

func newAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (s *AuthStorage) CreateUser(user models.User) error {
	records := `INSERT INTO users(username, email, password) VALUES (?,?,?)` // the query itself
	query, err := s.db.Prepare(records)                                      // prepares above-mentioned query INSERT INTO users(token) VALUES(?) WHERE users_id = ?
	if err != nil {
		return fmt.Errorf("an error occurred while preparing the file: %v", err)
	}
	_, err = query.Exec(user.Username, user.Email, user.Password) // executes this query
	if err != nil {
		return fmt.Errorf("an error occurred while executing the file: %v", err)
	}
	return nil
}

func (s *AuthStorage) DeleteSession(ID int64) error {
	_, err := s.db.Exec("DELETE FROM sessions WHERE user_id=?", ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthStorage) SaveSession(ID int64, token string, cookieTime time.Time) error {
	records := `INSERT INTO sessions(user_id,token,expiry) VALUES(?,?,?)` // the query itself
	query, err := s.db.Prepare(records)                                   // INSERT INTO users(token) VALUES(?) WHERE users_id = ?
	if err != nil {
		return err
	}
	_, err = query.Exec(ID, token, cookieTime) // executes this query
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthStorage) CheckIfTheSessionExists(ID int64) (bool, error) {
	find_like := s.db.QueryRow("SELECT session_id, user_id, token,expiry FROM sessions WHERE user_id", ID) // gets the user_id by token
	session := models.Session{}
	err := find_like.Scan(&session.Session_id, &session.User_id, &session.Token, &session.Expiry)
	if err != nil {
		return false, nil // the user doesnt have a session before
	}
	return true, nil
}

func (s *AuthStorage) GetUser(u models.User) (models.User, error) {
	row := s.db.QueryRow("SELECT * FROM users WHERE username = ?", u.Username) // gets the username from database
	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		fmt.Println(err)
		return models.User{}, err
	} // if u.Username does not exist, method Scan returns error! row returns the columns with matched username. If scan can not scan or find, then it returns error
	return user, nil
}

func (s *AuthStorage) CheckEmailforDuplicate(u models.User) error {
	row := s.db.QueryRow("SELECT * FROM users WHERE email = ?", u.Email) // gets the email from database
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password); err != nil {
		fmt.Println(err)
		return err
	} // if u.Username does not exist, method Scan returns error! row returns the columns with matched username. If scan can not scan or find, then it returns error
	return nil
}
