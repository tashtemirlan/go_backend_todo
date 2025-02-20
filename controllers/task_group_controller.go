package controllers

import (
	"github.com/gin-gonic/gin"
	"material_todo_go/database"
	"material_todo_go/models"
	"material_todo_go/utils"
	"net/http"
	"strings"
)

// CreateTaskGroup handles creating a new task group
func CreateTaskGroup(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	_, err := utils.ParseJWT(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var taskGroup models.TaskGroup
	if err := c.ShouldBindJSON(&taskGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	database.DB.Create(&taskGroup)
	c.JSON(http.StatusCreated, gin.H{})
}

// GetTaskGroups retrieves all task groups
func GetTaskGroups(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	_, err := utils.ParseJWT(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var taskGroups []models.TaskGroup
	database.DB.Find(&taskGroups)
	c.JSON(http.StatusOK, taskGroups)
}

// GetTaskGroup retrieves a single task group by ID
func GetTaskGroup(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	_, err := utils.ParseJWT(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var taskGroup models.TaskGroup
	if err := database.DB.First(&taskGroup, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task group not found"})
		return
	}

	c.JSON(http.StatusOK, taskGroup)
}

// UpdateTaskGroup updates an existing task group
func UpdateTaskGroup(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	_, err := utils.ParseJWT(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var taskGroup models.TaskGroup
	if err := database.DB.First(&taskGroup, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task group not found"})
		return
	}

	if err := c.ShouldBindJSON(&taskGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	database.DB.Save(&taskGroup)
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteTaskGroup deletes a task group
func DeleteTaskGroup(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	_, err := utils.ParseJWT(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if err := database.DB.Delete(&models.TaskGroup{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task group"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task group deleted successfully"})
}
