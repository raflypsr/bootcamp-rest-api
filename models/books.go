package models

import (
	"time"
)

type (
	Books struct {
		ID          uint      `gorm:"primary_key" json:"id"`
		Judul       string    `gorm:"type:varchar(255); not null" json:"judul"`
		Penulis     string    `gorm:"type:varchar(255); not null" json:"penulis"`
		Penerbit    string    `gorm:"type:varchar(255); not null" json:"penerbit"`
		TahunTerbit int       `gorm:"type:int; not null" json:"tahun_terbit"`
		CreatedAt   time.Time `gorm:"not null" json:"created_at"`
		UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
		Reviews     []Reviews `gorm:"foreignKey:BookID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	}
)
