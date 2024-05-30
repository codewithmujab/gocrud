package models

import (
	"gorm.io/gorm"
)

// book model
type Book struct {
	gorm.Model

	Title  string `json:"title" validate:"required,min=3,max=255"`
	Author string `json:"author" validate:"required,min=3,max=100"`
	//Email string `json:"email" validate:"required,email,min=3,max=100"`
}
