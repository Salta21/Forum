package storage

import (
	"database/sql"
	"fmt"
	"strconv"

	"forum/web/models"
)

type Post interface {
	GetPostsById(user_id int) ([]models.Post, error)
	SelectContentsByPostId(post_id int) ([]string, error)
	GetLikedPostsById(user_id int) ([]models.LikedPosts, error)
	IfUserLikedPost(user_id string, post_id string) (bool, error)
	DeleteLikedPost(usId, postId int) error
	IfUserDislikedPost(user_id string, post_id string) (bool, error)
	DeleteDislikedPost(usId, postId int) error
	FillTheDislikesPostTable(usId, postId int) error
	FillTheLikesPostTable(usId, postId int) error
}

type PostStorage struct {
	db *sql.DB
}

func newPostStorage(db *sql.DB) *PostStorage {
	return &PostStorage{
		db: db,
	}
}

func (s *PostStorage) GetPostsById(user_id int) ([]models.Post, error) { // you need to finish this function by getting all the posts by this id and display them
	rows, err := s.db.Query("SELECT posts_id,user_id,posts,author,title FROM posts WHERE user_id = ?", user_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	allPost := []models.Post{}
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Post_id, &post.User_id, &post.Posts, &post.Author, &post.Title); err != nil {
			return nil, err
		}
		contents, err := s.SelectContentsByPostId(post.Post_id)
		if err != nil {
			return nil, err
		}
		post.Contents = contents
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		allPost = append(allPost, post)
	}

	return allPost, nil
}

func (s *PostStorage) SelectContentsByPostId(post_id int) ([]string, error) {
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

func (s *PostStorage) GetLikedPostsById(user_id int) ([]models.LikedPosts, error) { // you need to finish this function by getting all the posts by this id and display them
	rows, err := s.db.Query("SELECT post_id FROM likes WHERE user_id = ?", user_id)
	if err != nil {
		return nil, err
	}
	allPost := []models.LikedPosts{}
	for rows.Next() {
		var post models.LikedPosts
		err := rows.Scan(&post.Post_id)
		if err != nil {
			return nil, err
		}
		allPost = append(allPost, post)
	}
	return allPost, nil
}

func (s *PostStorage) IfUserLikedPost(user_id string, post_id string) (bool, error) {
	usId, err := strconv.Atoi(user_id)
	if err != nil {
		return false, err
	}
	postId, err := strconv.Atoi(post_id)
	if err != nil {
		return false, err
	}
	find_like := s.db.QueryRow("SELECT user_id, post_id FROM likes WHERE post_id = ? AND user_id = ?", postId, usId) // gets the user_id by token
	like := models.LikedPosts{}
	err = find_like.Scan(&like.User_id, &like.Post_id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (s *PostStorage) DeleteLikedPost(usId, postId int) error {
	_, err := s.db.Exec("DELETE FROM likes WHERE user_id=? AND post_id=?", usId, postId)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStorage) IfUserDislikedPost(user_id string, post_id string) (bool, error) {
	usId, err := strconv.Atoi(user_id)
	if err != nil {
		return false, err
	}
	postId, err := strconv.Atoi(post_id)
	if err != nil {
		return false, err
	}
	find_like := s.db.QueryRow("SELECT user_id, post_id  FROM dislikes WHERE post_id = ? AND user_id = ?", postId, usId) // gets the user_id by token
	user := models.Post{}
	err = find_like.Scan(&user.User_id, &user.Post_id)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (s *PostStorage) DeleteDislikedPost(usId, postId int) error {
	_, err := s.db.Exec("DELETE FROM dislikes WHERE user_id=? AND post_id=?", usId, postId)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStorage) FillTheDislikesPostTable(usId, postId int) error {
	records := `INSERT INTO dislikes(user_id, post_id) VALUES (?,?)` // the query itself
	query, err := s.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(usId, postId) // executes this query
	if err != nil {
		return err
	}
	return nil
}

func (s *PostStorage) FillTheLikesPostTable(usId, postId int) error {
	records := `INSERT INTO likes(user_id, post_id) VALUES (?,?)` // the query itself
	query, err := s.db.Prepare(records)                           // prepares above-mentioned query INSERT INTO users(token) VALUES(?) WHERE users_id = ?
	if err != nil {
		return err
	}
	_, err = query.Exec(usId, postId) // executes this query
	if err != nil {
		return err
	}
	return nil
}
