package cmd

import (
	login "github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/models"
	user "github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/models"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(user.User{})
	db.AutoMigrate(login.Login{})
	// db.AutoMigrate()
}
