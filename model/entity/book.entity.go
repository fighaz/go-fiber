package entity

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement" `
	Title     string         `json:"title"`
	Author    string         `json:"author"  gorm:"size:255"`
	Cover     string         `json:"cover"  gorm:"size:255"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
