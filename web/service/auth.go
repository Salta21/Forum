package service

import (
	"errors"
	"net/http"
	"regexp"
	"time"
	"unicode"

	"forum/web/models"
	"forum/web/storage"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAuth              = errors.New("auth error")
	ErrUserExist         = errors.New("user exists already")
	ErrEmailInvalid      = errors.New("email invalid")
	ErrEmailExist        = errors.New("email exists already")
	ErrPasswordInvalid   = errors.New("password invalid")
	ErrUserNotExist      = errors.New("user does not exist (sign in)")
	ErrIncorrectPassword = errors.New("password incorrect")
)

type Auth interface {
	CreateUser(user models.User) error
	GenerateSessionToken(user models.User) (http.Cookie, error)
	DeleteSessionToken(user_id int64) error
}

type AuthService struct {
	storage storage.Auth
}

func newAuthService(storage storage.Auth) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) CreateUser(user models.User) error {
	checkUsername := IsUsernameValid(user.Username)
	if !checkUsername {
		return ErrAuth
	}

	_, err := s.storage.GetUser(user) // before sending to database, first, check if the user exists,
	if err == nil {                   // if there is no error in finding the matched username, so the user exists
		return ErrUserExist
	}

	check := IsEmailValid(user.Email)
	if !check {
		return ErrEmailInvalid
	}
	err = s.storage.CheckEmailforDuplicate(user)
	if err == nil { // if there is no error in finding the matched username, so the user exists
		return ErrEmailExist
	}
	checkPassword := IsPasswordValid(user.Password)
	if !checkPassword {
		return ErrPasswordInvalid
	}
	user.Password, err = HashThePassword(user.Password)
	if err != nil {
		return err
	}
	if err := s.storage.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) GenerateSessionToken(u models.User) (http.Cookie, error) {
	user, err := s.storage.GetUser(u) // before sending to database, first, check if the user exists,
	if err != nil {                   // if there is no error in finding the matched username, so the user exists
		return http.Cookie{}, ErrUserNotExist
	}

	if CompareHashPassword([]byte(user.Password), []byte(u.Password)) != nil {
		return http.Cookie{}, ErrIncorrectPassword
	}
	token, err := token() // token I took from another function
	if err != nil {
		return http.Cookie{}, err
	}
	cookie := http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = token
	cookie.Expires = time.Now().Add(15 * time.Minute)
	check, err := s.storage.CheckIfTheSessionExists(user.ID)
	if err != nil {
		return http.Cookie{}, err
	}
	if check {
		if err := s.storage.DeleteSession(user.ID); err != nil {
			return http.Cookie{}, err
		}
	}
	if err := s.storage.SaveSession(user.ID, cookie.Value, cookie.Expires); err != nil {
		return http.Cookie{}, err
	}
	return cookie, nil
}

func (s *AuthService) DeleteSessionToken(user_id int64) error {
	err := s.storage.DeleteSession(user_id)
	if err != nil {
		return err
	}
	return nil
}

func token() (string, error) {
	u2, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u2.String(), nil
}

func IsPasswordValid(pass string) bool { // check is the password valid
	var (
		upp, low, num, sym bool
		tot                uint8
	)
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}
	if !upp || !low || !num || !sym || tot < 6 {
		return false
	}
	return true
}

func IsUsernameValid(s string) bool { // check if the username valid
	usernameRegex := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]{5,30}$`)
	return usernameRegex.MatchString(s)
}

func IsEmailValid(s string) bool { // check if email is valid
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{1,30}$`)
	return emailRegex.MatchString(s)
}

func HashThePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHashPassword(user_password, password []byte) error {
	if err := bcrypt.CompareHashAndPassword(user_password, password); err != nil {
		return err
	}
	return nil
}
