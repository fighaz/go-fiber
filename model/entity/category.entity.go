package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement" `
	Name      string `json:"name"  gorm:"size:255"`
	Photos    []Photo
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
