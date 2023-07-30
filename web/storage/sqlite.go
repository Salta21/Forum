package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	usersTable = `CREATE TABLE IF NOT EXISTS users(
		users_id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		email TEXT,
		password TEXT
		);`
	sessionsTable = `CREATE TABLE IF NOT EXISTS sessions(
		session_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		token TEXT,
		expiry DATE
		);`
	postsTable = `CREATE TABLE IF NOT EXISTS posts(
		posts_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		posts TEXT,
		author TEXT,
		title TEXT
	  );`
	likesTable = `CREATE TABLE IF NOT EXISTS likes(
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NUll
	)`
	dislikesTable = `CREATE TABLE IF NOT EXISTS dislikes(
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NUll
	)`
	commentsTable = `CREATE TABLE IF NOT EXISTS comments(
		comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
		commenter_id INTEGER NOT NULL,
		post_id INTEGER NOT NUll,
		commenttext TEXT
	)`
	likecommentsTable = `CREATE TABLE IF NOT EXISTS likecomments(
		liker_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL
	)`
	dislikecommentsTable = `CREATE TABLE IF NOT EXISTS dislikecomments(
		disliker_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL
	)`
	postcontentTable = `CREATE TABLE IF NOT EXISTS postcontent(
		post_id INTEGER NOT NULL,
		content TEXT
	)`
)

func CreateDB() (*sql.DB, error) { // first sql1 table (id, username, email,password)
	db, err := sql.Open("sqlite3", "store.db?_foreign_keys=on")
	if err != nil {
		return nil, fmt.Errorf("storage: init db: %w", err)
	}
	return db, nil
}

func CreateTables(db *sql.DB) error {
	tables := []string{usersTable, sessionsTable, postsTable, likesTable, dislikesTable, commentsTable, likecommentsTable, dislikecommentsTable, postcontentTable}
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			return err
		}
	}
	return nil
}
