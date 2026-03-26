package models


import "time"

// User struct holds the data for a user
type User struct {
	ID       int    // User ID from database
	Username string // Name from register form
	Email    string // Email from register form
	Password string // Hashed password (not plain text)
}

// Session struct holds the data for a logged-in user
type Session struct {
	UUID      string    // The unique cookie value
	UserID    int       // ID of the user who owns this session
	ExpiresAt time.Time // When the cookie will stop working
}