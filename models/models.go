package models

import "time"

type User struct {
	ID        int
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	ID        string
	UserID    int
	ExpiresAt time.Time
}

type Post struct {
	ID        int
	UserID    int
	Username  string // joined from users
	Title     string
	Content   string
	CreatedAt time.Time
}

type Comment struct {
	ID        int
	PostID    int
	UserID    int
	Username  string // joined from users
	Content   string
	CreatedAt time.Time
}

type Category struct {
	ID   int
	Name string
}

type Like struct {
	ID        int
	UserID    int
	PostID    int
	CommentID int
	Value     int // +1 or -1
}