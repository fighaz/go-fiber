package entity

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID         uint           `json:"id" gorm:"primaryKey;autoIncrement" `
	Image      string         `json:"image"  gorm:"size:255"`
	CategoryId uint           `json:"category_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
