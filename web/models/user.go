package models

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	Erstring  string
	Post      []Post // this hould be checked one more time
	LikedPost []LikedPosts
	Contents  []string
}
