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

	// Define response struct
	type TaskGroupResponse struct {
		ID              uint   `json:"id"`
		Name            string `json:"name"`
		Description     string `json:"description"`
		IconData        int    `json:"icon_data"`
		BackgroundColor string `json:"background_color"`
		IconColor       string `json:"icon_color"`
		UserID          uint   `json:"user_id"`
		TotalTasks      int64  `json:"total_tasks"`
		CompletionRate  int    `json:"completion_rate"`
	}

	var response []TaskGroupResponse

	for _, group := range taskGroups {
		var totalTasks int64
		var completedTasks int64

		// Count total tasks for this task group
		database.DB.Model(&models.Task{}).Where("task_group_id = ?", group.ID).Count(&totalTasks)

		// Count completed tasks for this task group
		database.DB.Model(&models.Task{}).Where("task_group_id = ? AND status = ?", group.ID, "COMPLETED").Count(&completedTasks)

		// Calculate completion rate
		completionRate := 0
		if totalTasks > 0 {
			completionRate = int(float64(completedTasks) / float64(totalTasks) * 100)
		}

		// Add to response
		response = append(response, TaskGroupResponse{
			ID:              group.ID,
			Name:            group.Name,
			Description:     group.Description,
			IconData:        group.IconData,
			BackgroundColor: group.BackgroundColor,
			IconColor:       group.IconColor,
			UserID:          group.UserID,
			TotalTasks:      totalTasks,
			CompletionRate:  completionRate,
		})
	}

	c.JSON(http.StatusOK, response)
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

// GetTasksWithCompletionPercentage get percentage of how much done and tasks for this group
func GetTasksWithCompletionPercentage(c *gin.Context) {
	taskGroupID := c.Param("id")

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

	// Get all tasks for the given task group
	var tasks []models.Task
	if err := database.DB.Where("task_group_id = ?", taskGroupID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	// Calculate completion percentage
	totalTasks := len(tasks)
	completedTasks := 0
	for _, task := range tasks {
		if task.Status == "COMPLETED" {
			completedTasks++
		}
	}

	completionPercentage := 0.0
	if totalTasks > 0 {
		completionPercentage = (float64(completedTasks) / float64(totalTasks)) * 100
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks":                tasks,
		"completionPercentage": completionPercentage,
	})
}
