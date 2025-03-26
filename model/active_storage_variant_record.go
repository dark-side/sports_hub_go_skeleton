package model

import (
	"time"
)

type ActiveStorageVariantRecord struct {
	ID              uint `gorm:"primaryKey"`
	BlobID          uint
	Blob            ActiveStorageBlob `gorm:"foreignKey:BlobID"`
	VariationDigest string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
