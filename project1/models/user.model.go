package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey;AUTO_INCREMENT;not null"`
	Name      string         `json:"name" form:"name" validate:"gte=6,lte=32" gorm:"not null;type:varchar(255)"`
	Email     string         `json:"email" form:"email" validate:"required,email" gorm:"not null;unique_index; type:varchar(255)"`
	Password  string         `json:"-" form:"password" validate:"required,gte=8" gorm:"not null;colum:password;type:varchar(255)"`
	Phone     int            `json:"phone" form:"phone" validate:"required,number,min=12" gorm:"required;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
