package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	TaskGroupID uint           `json:"task_group_id"`
	StartDate   time.Time      `json:"start_date"`
	FinishDate  time.Time      `json:"finish_date"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
