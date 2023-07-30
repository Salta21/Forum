package models

type DisplayPost struct {
	Username      string
	Author        string
	Title         string
	Content       string
	Posts         string
	Countliked    int
	Countdisliked int
	Forcomment    []PostTheComment
}
