package models

import (
	"time"
)

type (
	Likes struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		CreatedAt time.Time `gorm:"not null" json:"created_at"`
		UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
		UserID    uint      `gorm:"not null; uniqueIndex:idx_like" json:"user_id"`
		ReviewID  uint      `gorm:"not null; uniqueIndex:idx_like" json:"review_id"`
	}
)
