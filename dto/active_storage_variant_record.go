package dto

import "time"

type ActiveStorageVariantRecordResponse struct {
	ID              uint      `json:"id"`
	BlobID          uint      `json:"blob_id"`
	VariationDigest string    `json:"variation_digest"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
