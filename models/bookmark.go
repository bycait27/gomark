package models

import (
	"time"

	"gorm.io/gorm"
)

type Bookmark struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	URL         string         `json:"url" gorm:"not null"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
