package repository

import (
	"go_be_plgrnd/model"

	"gorm.io/gorm"
)

type ActiveStorageBlobRepository interface {
	InsertBlob(blob model.ActiveStorageBlob) model.ActiveStorageBlob
	UpdateBlob(blob model.ActiveStorageBlob) model.ActiveStorageBlob
	DeleteBlob(blob model.ActiveStorageBlob) error
	FindBlobByID(blobID uint) model.ActiveStorageBlob
	FindBlobByKey(key string) model.ActiveStorageBlob
}

type blobConnection struct {
	connection *gorm.DB
}

func NewActiveStorageBlobRepository(db *gorm.DB) ActiveStorageBlobRepository {
	return &blobConnection{
		connection: db,
	}
}

func (db *blobConnection) InsertBlob(blob model.ActiveStorageBlob) model.ActiveStorageBlob {
	db.connection.Create(&blob)
	return blob
}

func (db *blobConnection) UpdateBlob(blob model.ActiveStorageBlob) model.ActiveStorageBlob {
	db.connection.Save(&blob)
	return blob
}

func (db *blobConnection) DeleteBlob(blob model.ActiveStorageBlob) error {
	return db.connection.Delete(&blob).Error
}

func (db *blobConnection) FindBlobByID(blobID uint) model.ActiveStorageBlob {
	var blob model.ActiveStorageBlob
	db.connection.First(&blob, blobID)
	return blob
}

func (db *blobConnection) FindBlobByKey(key string) model.ActiveStorageBlob {
	var blob model.ActiveStorageBlob
	db.connection.Where("key = ?", key).First(&blob)
	return blob
}
