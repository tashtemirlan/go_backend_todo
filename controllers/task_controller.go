package controllers

import (
	"material_todo_go/database"
	"material_todo_go/models"
	"material_todo_go/utils"
	"net/http"
	"strings"
	_ "time"

	"github.com/gin-gonic/gin"
)

// Extract and validate JWT token
func getUserFromToken(c *gin.Context) (string, bool) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return "", false
	}

	// Ensure "Bearer <token>" format
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return "", false
	}

	tokenString := parts[1]
	email, err := utils.ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return "", false
	}

	return email, true
}

// CreateTask creates a new task
func CreateTask(c *gin.Context) {
	_, valid := getUserFromToken(c)
	if !valid {
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// GetAllTasks returns all tasks
func GetAllTasks(c *gin.Context) {
	_, valid := getUserFromToken(c)
	if !valid {
		return
	}

	var tasks []models.Task
	database.DB.Find(&tasks)

	// Return empty array if no tasks found
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, []models.Task{})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTask retrieves a task by ID
func GetTask(c *gin.Context) {
	_, valid := getUserFromToken(c)
	if !valid {
		return
	}

	taskID := c.Param("id")
	var task models.Task

	if err := database.DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask updates an existing task
func UpdateTask(c *gin.Context) {
	_, valid := getUserFromToken(c)
	if !valid {
		return
	}

	taskID := c.Param("id")
	var task models.Task

	if err := database.DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var updatedData models.Task
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Update fields
	if updatedData.Title != "" {
		task.Title = updatedData.Title
	}
	if updatedData.Description != "" {
		task.Description = updatedData.Description
	}
	if updatedData.Status != "" {
		task.Status = updatedData.Status
	}
	if updatedData.TaskGroupID != 0 {
		task.TaskGroupID = updatedData.TaskGroupID
	}
	if !updatedData.StartDate.IsZero() {
		task.StartDate = updatedData.StartDate
	}
	if !updatedData.FinishDate.IsZero() {
		task.FinishDate = updatedData.FinishDate
	}

	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// DeleteTask deletes a task by ID
func DeleteTask(c *gin.Context) {
	_, valid := getUserFromToken(c)
	if !valid {
		return
	}

	taskID := c.Param("id")
	if err := database.DB.Delete(&models.Task{}, taskID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
