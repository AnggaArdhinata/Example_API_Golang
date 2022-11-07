package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   	`gorm:"primaryKey" json:"id"`
	Name      string 	`gorm:"type:varchar(255)" json:"name"`
	Email     string 	`gorm:"type:varchar(255)" json:"email"`
	Password  string 	`gorm:"type:varchar(255)" json:"password"`
	Role      string 	`gorm:"type:varchar(255)" json:"role"`
	CreatedAt time.Time	`gorm:"type:timestamp;default:now()" json:"creatd_at"`
	UpdatedAt time.Time	`gorm:"type:timestamp" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp" json:"deleted_at"`
}

type Users []User