package database

import (
	"bsku001/backend/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.MerchantFeatureConfig{},
		&entity.Brand{},
		&entity.Category{},
		&entity.SubCategory{},
		&entity.Material{},
		&entity.Color{},
		&entity.SKU{},
		&entity.SKUDimension{},
		&entity.SKUImage{},
		&entity.SKUOptionGroup{},
		&entity.SKUOptionValue{},
	)
}
