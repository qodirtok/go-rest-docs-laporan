package repositories

import (
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/models"
	"gorm.io/gorm"
)

type LoginRepositories interface {
	Save(login models.Login) (models.Login, error)
	FindByUsername(username string) (models.Login, error)
}

type repository struct {
	db *gorm.DB
}

func NewLoginRepositories(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(login models.Login) (models.Login, error) {
	err := r.db.Create(&login).Error
	// fmt.Println(err)
	return login, err
}

func (r *repository) FindByUsername(username string) (models.Login, error) {
	var login models.Login
	err := r.db.Where("username =?", username).Find(&login).Error

	return login, err
}
