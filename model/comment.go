package model

import (
	"time"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Text      string `gorm:"type:text"`
	UserID    uint
	User      User `gorm:"foreignKey:UserID"`
	ArticleID uint
	Article   Article `gorm:"foreignKey:ArticleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
