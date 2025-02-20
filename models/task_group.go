package models

import (
	"gorm.io/gorm"
)

// TaskGroup represents a group of tasks categorized by an icon and colors.
type TaskGroup struct {
	gorm.Model
	Name            string `json:"name" gorm:"not null"`
	IconData        int    `json:"icon_data" gorm:"not null"`
	BackgroundColor string `json:"background_color" gorm:"not null"`
	IconColor       string `json:"icon_color" gorm:"not null"`
	UserID          uint   `json:"user_id" gorm:"not null"` // Associate with a user
}
