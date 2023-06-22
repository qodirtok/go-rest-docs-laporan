package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255;not null"`
	Password string `gorm:"size:255"`
	Username string `gorm:"size:255;not null;unique"`
}

type UserInput struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
