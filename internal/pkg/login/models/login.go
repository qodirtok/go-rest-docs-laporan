package models

import "gorm.io/gorm"

type Login struct {
	*gorm.Model
	Username string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null"`
}

type AuthentificationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
