package main

import (
	"blog-backend/handlers"
    "blog-backend/models"
	"blog-backend/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var err error

	// Initialize SQLite database connection
	models.DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Automatically migrate models to create database tables
	models.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	// Initialize Gin router with default middleware (Logger and Recovery)
	r := gin.Default()

// Public routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/posts", handlers.GetPosts)                 // Get all posts
	r.GET("/posts/:id", handlers.GetPostByID)          // Get single post
	r.GET("/posts/:id/comments", handlers.GetCommentsByPost) // Get comments for a post

	// Protected routes (Requires JWT authentication)
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware())
	{
		authGroup.POST("/posts", handlers.CreatePost)          // Create post
		authGroup.PUT("/posts/:id", handlers.UpdatePost)       // Update post (Author only)
		authGroup.DELETE("/posts/:id", handlers.DeletePost)    // Delete post (Author only)
		authGroup.POST("/comments", handlers.CreateComment)    // Create comment
	}

	// Start the HTTP server on port 8080
	r.Run(":8080")
}
