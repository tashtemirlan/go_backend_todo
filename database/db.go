package database

import (
	"fmt"
	"log"
	"material_todo_go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"material_todo_go/config"
)

var DB *gorm.DB

func ConnectDB() {
	config.LoadConfig()

	// Print config values to debug
	fmt.Printf("Connecting to DB: host=%s port=%s user=%s dbname=%s sslmode=disable\n",
		config.DBHost, config.DBPort, config.DBUser, config.DBName)

	// Correct connection string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL!")
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Note{})
	DB.AutoMigrate(&models.TaskGroup{})
	DB.AutoMigrate(&models.Task{})
}
