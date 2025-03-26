package dto

import "time"

type ActiveStorageBlobResponse struct {
	ID          uint      `json:"id"`
	Key         string    `json:"key"`
	Filename    string    `json:"filename"`
	ContentType string    `json:"content_type"`
	Metadata    string    `json:"metadata"`
	ServiceName string    `json:"service_name"`
	ByteSize    uint      `json:"byte_size"`
	Checksum    string    `json:"checksum"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
