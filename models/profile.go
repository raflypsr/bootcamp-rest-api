package models

import (
	"time"
)

type (
	Profile struct {
		ID           uint      `gorm:"primary_key" json:"id"`
		Umur         int       `gorm:"type:int; not null" json:"umur"`
		NamaLengkap  string    `gorm:"type:varchar(255); not null" json:"nama_lengkap"`
		JenisKelamin string    `gorm:"type:varchar(255); not null" json:"jenis_kelamin"`
		CreatedAt    time.Time `gorm:"not null" json:"created_at"`
		UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`
		UserID       uint      `gorm:"unique; not null" json:"user_id"`
	}
)
