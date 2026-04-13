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
	Likes     int
	Dislikes  int
	CreatedAt time.Time
	Category string // joined from categories 
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

// PageData is passed to the index.html template, bundling posts with auth state.
type PageData struct {
	Posts      []Post
	IsLoggedIn bool
	Username   string
}

type PostPageData struct {
	Post       Post
	Comments   []Comment
	IsLoggedIn bool
	Username   string
}
