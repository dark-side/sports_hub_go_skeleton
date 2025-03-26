package dto

import "time"

type ActiveStorageAttachmentResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	RecordType string    `json:"record_type"`
	RecordID   uint      `json:"record_id"`
	BlobID     uint      `json:"blob_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
