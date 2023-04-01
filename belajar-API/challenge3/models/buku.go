package models

import "time"

type Buku struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	JudulBuku string    `gorm:"not null;type:varchar(200)" json:"judul_buku"`
	Author    string    `gorm:"not null;type:varchar(200)" json:"author"`
	Kategori  string    `gorm:"not null;type:varchar(200)" json:"kategori"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
