package handlers

import (
	"net/http"

	"blog-backend/models"

	"github.com/gin-gonic/gin"
)

// CreateComment allows an authenticated user to add a comment to a post.
func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get current logged-in user ID from JWT middleware and assign to comment
	comment.UserID = c.GetUint("userID")

	// Save comment to database
	if err := models.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully", "comment": comment})
}

// GetCommentsByPost retrieves all comments for a specific post.
func GetCommentsByPost(c *gin.Context) {
	postID := c.Param("id")
	var comments []models.Comment

	// Find comments belonging to the specific post and preload User info
	if err := models.DB.Where("post_id = ?", postID).Preload("User").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
