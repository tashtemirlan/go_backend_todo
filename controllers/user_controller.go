package controllers

import (
	"github.com/gin-gonic/gin"
	"material_todo_go/database"
	"material_todo_go/models"
	"material_todo_go/utils"
	"net/http"
	_ "os"
	"path/filepath"
	_ "path/filepath"
	"strings"
)

func GetUserInformation(c *gin.Context) {
	// Extract the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	// Split "Bearer <TOKEN>" and extract only the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	tokenString := parts[1] // Extract the actual JWT

	// Parse JWT token and get email
	email, err := utils.ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Find user by email
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Return user data (excluding password)
	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"full_name": user.FullName,
		"email":     user.Email,
		"image":     user.Image,
	})
}

func UpdateUser(c *gin.Context) {
	// Extract Authorization token
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	// Split "Bearer <TOKEN>" and extract only the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	tokenString := parts[1] // Extract the actual JWT

	// Parse JWT and get email
	email, err := utils.ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Find user by email
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Define variables for optional fields
	var newFullName string
	var filePath string

	// Detect if the request is JSON
	contentType := c.GetHeader("Content-Type")

	if strings.Contains(contentType, "application/json") {
		// Handle JSON request (for Flutter)
		var requestData struct {
			FullName string `json:"full_name"`
		}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			return
		}
		newFullName = requestData.FullName
	} else {
		// Handle form-data request (for Postman)
		newFullName = c.PostForm("full_name")

		// Handle image upload (if provided)
		file, err := c.FormFile("image")
		if err == nil {
			// Save uploaded image
			filePath = filepath.Join("uploads", file.Filename)

			// Convert backslashes to forward slashes for proper URL formatting
			filePath = strings.Replace(filePath, "\\", "/", -1)

			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
				return
			}
			user.Image = filePath // Update image only if a new one is uploaded
		}
	}

	// Update full_name if provided
	if newFullName != "" {
		user.FullName = newFullName
	}

	// Save changes to the database
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	// Return updated user data
	c.JSON(http.StatusOK, gin.H{})
}
