package repository

import (
	"go_be_plgrnd/model"

	"gorm.io/gorm"
)

type ActiveStorageVariantRecordRepository interface {
	InsertVariant(variant model.ActiveStorageVariantRecord) model.ActiveStorageVariantRecord
	UpdateVariant(variant model.ActiveStorageVariantRecord) model.ActiveStorageVariantRecord
	DeleteVariant(variant model.ActiveStorageVariantRecord) error
	FindVariantByID(variantID uint) model.ActiveStorageVariantRecord
	FindVariantsByBlobID(blobID uint) []model.ActiveStorageVariantRecord
}

type variantConnection struct {
	connection *gorm.DB
}

func NewActiveStorageVariantRecordRepository(db *gorm.DB) ActiveStorageVariantRecordRepository {
	return &variantConnection{
		connection: db,
	}
}

func (db *variantConnection) InsertVariant(variant model.ActiveStorageVariantRecord) model.ActiveStorageVariantRecord {
	db.connection.Create(&variant)
	return variant
}

func (db *variantConnection) UpdateVariant(variant model.ActiveStorageVariantRecord) model.ActiveStorageVariantRecord {
	db.connection.Save(&variant)
	return variant
}

func (db *variantConnection) DeleteVariant(variant model.ActiveStorageVariantRecord) error {
	return db.connection.Delete(&variant).Error
}

func (db *variantConnection) FindVariantByID(variantID uint) model.ActiveStorageVariantRecord {
	var variant model.ActiveStorageVariantRecord
	db.connection.First(&variant, variantID)
	return variant
}

func (db *variantConnection) FindVariantsByBlobID(blobID uint) []model.ActiveStorageVariantRecord {
	var variants []model.ActiveStorageVariantRecord
	db.connection.Where("blob_id = ?", blobID).Find(&variants)
	return variants
}
