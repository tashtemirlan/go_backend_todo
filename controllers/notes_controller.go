package controllers

import (
	"github.com/gin-gonic/gin"
	"material_todo_go/database"
	"material_todo_go/models"
	"net/http"
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

	c.JSON(http.StatusCreated, gin.H{})
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

	c.JSON(http.StatusOK, gin.H{})
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
