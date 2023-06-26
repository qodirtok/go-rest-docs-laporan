package models

import "gorm.io/gorm"

type Dokumen struct {
	*gorm.Model
	Title     string `gorm:"size:255"`
	DocsNumer string `gorm:"size:255"`
	Shipper   string `gorm:"size:255"`
	Reciver   string `gorm:"size:255"`
	Abuot     string `gorm:"size:255"`
	File      string `gorm:"size:255"`
}

type DokumenInput struct {
	Title     string `json:"title"`
	DocsNumer string `json:"docs_number"`
	Shipper   string `json:"shipper"`
	Reciver   string `json:"reciver"`
	Abuot     string `json:"about"`
	File      string `json:"file"`
}

type DokumenResponse struct {
	Title     string `json:"title"`
	DocsNumer string `json:"docs_number"`
	Shipper   string `json:"shipper"`
	Reciver   string `json:"reciver"`
	Abuot     string `json:"about"`
	File      string `json:"file"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
