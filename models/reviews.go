package models

import (
	"time"
)

type (
	Reviews struct {
		ID        uint       `gorm:"primary_key" json:"id"`
		Judul 	  string     `json:"judul"`
		Deskripsi string     `gorm:"not null" json:"deskripsi"`
		CreatedAt time.Time  `gorm:"not null" json:"created_at"`
		UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
		UserID    uint       `gorm:"not null; uniqueIndex:idx_review" json:"user_id"`
		BookID    uint       `gorm:"not null; uniqueIndex:idx_review" json:"book_id"`
		Likes     []Likes    `gorm:"foreignKey:ReviewID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
		Comments  []Comments `gorm:"foreignKey:ReviewID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	}
)
