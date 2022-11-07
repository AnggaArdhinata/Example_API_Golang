package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint		`gorm:"primaryKey" json:"id"`
	Name        string		`gorm:"type:varchar(255)" json:"name"`
	Price       string		`gorm:"type:varchar(255)" json:"price"`
	Description string		`gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time	`gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"type:timestamp" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp" json:"deleted_at"`
}

type Products []Product