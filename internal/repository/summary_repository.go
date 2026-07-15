package repository

import (
	"bsku001/backend/internal/entity"
	"bsku001/backend/internal/model"

	"gorm.io/gorm"
)

type SummaryRepository struct{ db *gorm.DB }

func NewSummaryRepository(db *gorm.DB) *SummaryRepository { return &SummaryRepository{db: db} }

func (r *SummaryRepository) Get(merchantID string) (model.SummaryResponse, error) {
	var summary model.SummaryResponse
	if err := r.db.Model(&entity.SKU{}).Where("merchant_id = ?", merchantID).Count(&summary.SumSKUs).Error; err != nil {
		return summary, err
	}
	if err := r.db.Model(&entity.Brand{}).Where("merchant_id = ?", merchantID).Count(&summary.SumBrand).Error; err != nil {
		return summary, err
	}
	if err := r.db.Model(&entity.Category{}).Where("merchant_id = ?", merchantID).Count(&summary.SumCategories).Error; err != nil {
		return summary, err
	}
	if err := r.db.Model(&entity.SubCategory{}).Where("merchant_id = ?", merchantID).Count(&summary.SumSubCategories).Error; err != nil {
		return summary, err
	}
	if err := r.db.Model(&entity.Material{}).Where("merchant_id = ?", merchantID).Count(&summary.SumMaterials).Error; err != nil {
		return summary, err
	}
	if err := r.db.Model(&entity.Color{}).Where("merchant_id = ?", merchantID).Count(&summary.SumColors).Error; err != nil {
		return summary, err
	}
	if err := r.db.Model(&entity.SKUOptionGroup{}).Where("merchant_id = ?", merchantID).Count(&summary.SumOptions).Error; err != nil {
		return summary, err
	}
	return summary, nil
}
