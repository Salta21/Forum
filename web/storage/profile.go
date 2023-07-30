package storage

import (
	"database/sql"
	"fmt"
	"log"

	"forum/web/models"
)

type Profile interface {
	GetUserById(n int) (string, string, error)
	GetAllPost() ([]models.Post, error)
	SelectContentsByPostId(post_id int) ([]string, error)
	FillThePostTable(userId int, post string, username string, title string) error
	FindPostIdbyPost(posts string) (int, error)
	FillTheContentTable(post_id int, content string) error
	GetAllPostsByCategory(c string) ([]models.Post, error)
	GetPostByPostId(postId int) (models.Post, error)
}
type ProfileStorage struct {
	db *sql.DB
}

func newProfileStorage(db *sql.DB) *ProfileStorage {
	return &ProfileStorage{
		db: db,
	}
}

func (s *ProfileStorage) GetUserById(n int) (string, string, error) {
	findId := s.db.QueryRow("SELECT * FROM users WHERE  users_id = ?", n) // gets the user_id by token
	user := models.User{}
	err := findId.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return "", "", err
	}
	return user.Username, user.Email, nil
}

func (s *ProfileStorage) GetAllPost() ([]models.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	allPost := []models.Post{}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Post_id, &post.User_id, &post.Posts, &post.Author, &post.Title)
		contents, _ := s.SelectContentsByPostId(post.Post_id)
		post.Contents = contents
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		allPost = append(allPost, post)
	}
	return allPost, nil
}

func (s *ProfileStorage) SelectContentsByPostId(post_id int) ([]string, error) {
	var check []string // it is empty
	rows, err := s.db.Query("SELECT content FROM postcontent WHERE post_id = ?", post_id)
	if err != nil {
		return check, err
	}
	var allPost []string
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Content)
		if err != nil {
			return check, err
		}
		allPost = append(allPost, post.Content)
	}
	return allPost, nil
}

func (s *ProfileStorage) FillThePostTable(userId int, post string, username string, title string) error {
	records := `INSERT INTO posts(user_id, posts, author, title) VALUES (?,?,?,?)` // the query itself
	query, err := s.db.Prepare(records)                                            // prepares above-mentioned query INSERT INTO users(token) VALUES(?) WHERE users_id = ?
	if err != nil {
		return err
	}
	_, err = query.Exec(userId, post, username, title) // executes this query
	if err != nil {
		return err
	}
	return nil
}

func (s *ProfileStorage) FindPostIdbyPost(posts string) (int, error) {
	findId := s.db.QueryRow("SELECT posts_id FROM posts WHERE  posts = ?", posts) // gets the user_id by token
	session := models.Post{}
	err := findId.Scan(&session.Post_id)
	if err != nil {
		return 0, err
	}
	return session.Post_id, nil
}

func (s *ProfileStorage) FillTheContentTable(post_id int, content string) error {
	records := `INSERT INTO postcontent(post_id, content) VALUES (?,?)`
	query, err := s.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(post_id, content) // executes this query
	if err != nil {
		return err
	}
	return nil
}

func (s *ProfileStorage) GetAllPostsByCategory(c string) ([]models.Post, error) {
	rows, err := s.db.Query("SELECT * FROM postcontent WHERE content = ?", c)
	if err != nil {
		return nil, err
	}
	allPost := []models.Post{}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Post_id, &post.Content)
		post, _ = s.GetPostByPostId(post.Post_id)
		if err != nil {
			return nil, err
		}
		allPost = append(allPost, post)
	}
	return allPost, nil
}

func (s *ProfileStorage) GetPostByPostId(postId int) (models.Post, error) {
	rows, err := s.db.Query("SELECT user_id, posts_id,author, title,posts FROM posts WHERE posts_id = ?", postId)
	if err != nil {
		return models.Post{}, err
	}
	var post models.Post
	for rows.Next() {
		err := rows.Scan(&post.User_id, &post.Post_id, &post.Author, &post.Title, &post.Posts)
		if err != nil {
			return models.Post{}, err
		}
		contents, err := s.SelectContentsByPostId(post.Post_id)
		if err != nil {
			return models.Post{}, err
		}
		post.Contents = contents
	}
	return post, nil
}
