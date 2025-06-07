package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement" `
	Name      string         `json:"name"  gorm:"size:255"`
	Email     string         `json:"email" gorm:"unique;size:255"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone" gorm:"size:15"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
