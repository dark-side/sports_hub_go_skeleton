package model

import (
	"time"
)

type Article struct {
	ID               uint `gorm:"primaryKey"`
	Title            string
	ShortDescription string
	Description      string `gorm:"type:text"`
	UserID           uint64
	User             User `gorm:"foreignKey:UserID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
