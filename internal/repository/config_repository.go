package repository

import (
	"errors"

	"bsku001/backend/internal/entity"

	"gorm.io/gorm"
)

type ConfigRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) *ConfigRepository { return &ConfigRepository{db: db} }

func (r *ConfigRepository) GetMerchantFeatureConfig(merchantID, userID string) (entity.MerchantFeatureConfig, error) {
	var item entity.MerchantFeatureConfig
	err := r.db.Where("merchant_id = ?", merchantID).First(&item).Error
	if err == nil {
		return item, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.MerchantFeatureConfig{}, err
	}
	item = entity.MerchantFeatureConfig{
		MerchantID:    merchantID,
		SKUs:          true,
		Brands:        true,
		Categories:    true,
		SubCategories: true,
		Materials:     true,
		Dimensions:    true,
		Images:        true,
		Colors:        true,
		SKUOptions:    true,
		CreatedBy:     userID,
		UpdatedBy:     userID,
	}
	return item, r.db.Create(&item).Error
}

func (r *ConfigRepository) UpdateMerchantFeatureConfig(item entity.MerchantFeatureConfig) (entity.MerchantFeatureConfig, error) {
	if err := r.db.Save(&item).Error; err != nil {
		return entity.MerchantFeatureConfig{}, err
	}
	return item, nil
}
