package repositories

import (
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/models"
	"gorm.io/gorm"
)

type UserRepositories interface {
	Create(user models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindOne(ID int) (models.User, error)
	Delete(ID int) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepositories(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) FindAll() ([]models.User, error) {
	var user []models.User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *repository) FindOne(ID int) (models.User, error) {
	var user models.User
	err := r.db.Where("id =?", ID).Find(&user).Error
	return user, err
}

func (r *repository) Delete(ID int) (models.User, error) {
	var user models.User
	err := r.db.Where("id =?", ID).Delete(&user).Error
	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}
