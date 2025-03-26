package model

import (
	"time"
)

type ActiveStorageBlob struct {
	ID          uint   `gorm:"primaryKey"`
	Key         string `gorm:"uniqueIndex"`
	Filename    string
	ContentType string
	Metadata    string `gorm:"type:text"`
	ServiceName string
	ByteSize    uint
	Checksum    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
