package repositories

import (
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/dokumen/models"
	"gorm.io/gorm"
)

type DokumenRepositories interface {
}

type repository struct {
	db *gorm.DB
}

func NewDokumenRepositories(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(dokumen models.Dokumen) (models.Dokumen, error) {
	err := r.db.Create(&dokumen).Error
	return dokumen, err
}

func (r *repository) FindAll(dokumen models.Dokumen) (models.Dokumen, error) {

}
