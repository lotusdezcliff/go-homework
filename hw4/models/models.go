package models

import (
	"gorm.io/gorm"
)

// DB is the global database instance used across the application.
var DB *gorm.DB

// User model represents a registered user in the system.
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

// Post model represents a blog article created by a user.
type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

// Comment model represents a user's comment on a specific post.
type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}
