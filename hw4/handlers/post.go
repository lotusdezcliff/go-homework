package handlers

import (
	"net/http"

	"blog-backend/models"

	"github.com/gin-gonic/gin"
)

// CreatePost allows an authenticated user to create a new blog post.
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from JWT middleware (stored in context)
	if err := models.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully", "post": post})
}

// GetPosts retrieves a list of all blog posts.
func GetPosts(c *gin.Context) {
	var posts []models.Post
	// Preload User information so we know who wrote the post
	if err := models.DB.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPostByID retrieves a single blog post by its ID.
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := models.DB.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost allows only the author to update their post.
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	// 1. Check if the post exists
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// 2. Get current logged-in user ID from JWT middleware context
	currentUserID := c.GetUint("userID")

	// 3. Verify if the current user is the author of the post
	if post.UserID != currentUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to update this post"})
		return
	}

	// 4. Bind new data and update
	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.Title = input.Title
	post.Content = input.Content
	models.DB.Save(&post)

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}

// DeletePost allows only the author to delete their post.
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	// 1. Check if the post exists
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// 2. Get current logged-in user ID from JWT middleware context
	currentUserID := c.GetUint("userID")

	// 3. Verify if the current user is the author of the post
	if post.UserID != currentUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to delete this post"})
		return
	}

	// 4. Delete the post
	models.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
