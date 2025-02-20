package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"material_todo_go/database"
	"material_todo_go/models"
	"material_todo_go/utils"
	"math/rand"
	"net/http"
	_ "os"
	"path/filepath"
	"strings"
	"time"
)

// @Summary Вход пользователя
// @Description Logs in a user and returns JWT
// @Tags Auth
// @Accept  json
// @Accept  multipart/form-data
// @Produce  json
// @Param email formData string false "Email"
// @Param password formData string false "Password"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Router /login [post]
func Login(c *gin.Context) {
	var request map[string]string

	// Check if request is JSON
	if err := c.ShouldBindJSON(&request); err != nil {
		// If JSON binding fails, try reading form-data
		email := c.PostForm("email")
		password := c.PostForm("password")

		// Ensure both fields are provided
		if email == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
			return
		}

		// Assign extracted values to request map
		request = map[string]string{
			"email":    email,
			"password": password,
		}
	}

	// Validate input
	if request["email"] == "" || request["password"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	var user models.User
	database.DB.Where("email = ?", request["email"]).First(&user)

	// Check if user exists and password is correct
	if user.ID == 0 || !utils.CheckPasswordHash(request["password"], user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, _ := utils.GenerateJWT(user.Email)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary Регистрация пользователя
// @Description Registers a new user
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param full_name formData string true "Full Name"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param image formData file false "Profile Image"  // Make the image parameter optional
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /signup [post]
func Signup(c *gin.Context) {
	var request struct {
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var fullName, email, password string

	// Try JSON binding first
	if err := c.ShouldBindJSON(&request); err == nil {
		fullName = request.FullName
		email = request.Email
		password = request.Password
	} else {
		// If JSON binding fails, fallback to form-data
		fullName = c.PostForm("full_name")
		email = c.PostForm("email")
		password = c.PostForm("password")
	}

	// Validate required fields
	if fullName == "" || email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Full Name, Email, and Password are required"})
		return
	}

	// Default image if not uploaded
	defaultImagePath := "uploads/default_avatar.png"
	var filePath string

	// Handle file upload
	file, err := c.FormFile("image")
	if err == nil {
		// Save the file
		filePath = filepath.Join("uploads", file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
	} else {
		// Use default image if no file was uploaded
		filePath = defaultImagePath
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user struct
	user := models.User{
		FullName: fullName,
		Email:    email,
		Password: hashedPassword,
		Image:    filePath,
	}

	// Save user in the database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return the created user (excluding password)
	c.JSON(http.StatusCreated, gin.H{})
}

// GenerateRandomCode generates a 6-digit random code
func GenerateRandomCode() string {
	rand.Seed(time.Now().UnixNano()) // Ensure randomness
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// @Summary Send password reset code
// @Description Sends a 6-digit reset code to the user's email.
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "User's email"
// @Router /forgetpassword/sendCode [post]
func SendResetCode(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil || request.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	// Check if user exists
	var user models.User
	if err := database.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Generate a 6-digit reset code
	resetCode := GenerateRandomCode()

	// Store the reset code in the database (optional)
	// You may store this in a `password_resets` table if needed.

	c.JSON(http.StatusOK, gin.H{"code": resetCode})
}

// @Summary Reset user password
// @Description Resets the user's password using only email and new password (no code required).
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "User's email"
// @Param password formData string true "New password"
// @Router /forgetpassword/resetPassword [post]
func ResetPassword(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil || request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if user exists
	var user models.User
	if err := database.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Update password
	if err := database.DB.Model(&models.User{}).Where("email = ?", request.Email).Update("password", hashedPassword).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
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
