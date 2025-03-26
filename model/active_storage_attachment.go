package model

import (
	"time"
)

type ActiveStorageAttachment struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"unique"`
	RecordType string
	RecordID   uint
	BlobID     uint
	Blob       ActiveStorageBlob `gorm:"foreignKey:BlobID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
