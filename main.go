package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	_ "material_todo_go/config"
	"material_todo_go/database"
	_ "material_todo_go/docs" // Import for Swagger
	"material_todo_go/routes"
)

// @title Material To-Do API
// @version 1.0
// @description Backend API for Material To-Do App
// @host localhost:2525
// @BasePath /api

func main() {
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	// Initialize database
	database.ConnectDB()

	// Setup routes
	routes.SetupRoutes(r)

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server running on port 2525")
	log.Println("http://localhost:2525/swagger/index.html")
	r.Run(":2525")
}
