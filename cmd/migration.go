package cmd

import (
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/models"
	"gorm.io/gorm"
)

func migration(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
