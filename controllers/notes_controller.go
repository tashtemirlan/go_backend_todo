package controllers

import (
	"errors"
	"material_todo_go/database"
	"material_todo_go/models"
	"material_todo_go/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CreateNote - Adds a new note for the authenticated user
func CreateNote(c *gin.Context) {
	userID, err := getAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	note.UserID = userID
	if err := database.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusCreated, note)
}

// GetAllNotes - Retrieves all notes for the authenticated user
func GetAllNotes(c *gin.Context) {
	userID, err := getAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var notes []models.Note
	result := database.DB.Where("user_id = ?", userID).Find(&notes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve notes"})
		return
	}

	// Always return an empty array if there are no notes
	if len(notes) == 0 {
		c.JSON(http.StatusOK, []models.Note{})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// GetNoteByID - Retrieves a single note by ID for the authenticated user
func GetNoteByID(c *gin.Context) {
	userID, err := getAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, note)
}

// UpdateNote - Updates an existing note for the authenticated user
func UpdateNote(c *gin.Context) {
	userID, err := getAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var updateData struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	note.Title = updateData.Title
	note.Description = updateData.Description

	if err := database.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, note)
}

// DeleteNote - Deletes a note for the authenticated user
func DeleteNote(c *gin.Context) {
	userID, err := getAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).Delete(&models.Note{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

// getAuthenticatedUserID extracts user ID from the JWT token
func getAuthenticatedUserID(c *gin.Context) (uint, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return 0, errors.New("Authorization token required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, errors.New("Invalid Authorization header format")
	}

	email, err := utils.ParseJWT(parts[1])
	if err != nil {
		return 0, errors.New("Invalid token")
	}

	// Retrieve user ID based on email
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return 0, errors.New("User not found")
	}

	return user.ID, nil
}
