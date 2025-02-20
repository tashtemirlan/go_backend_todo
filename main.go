package main

import (
	"github.com/gin-gonic/gin"
	"log"
	_ "material_todo_go/config"
	"material_todo_go/database"
	"material_todo_go/routes"
)

func main() {
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	// Initialize database
	database.ConnectDB()

	// Setup routes
	routes.SetupRoutes(r)

	log.Println("Server running on port 2525")
	r.Run(":2525")
}
