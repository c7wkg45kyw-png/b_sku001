package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantFeatureConfig struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	MerchantID    string    `gorm:"size:80;not null;uniqueIndex" json:"merchant_id"`
	SKUs          bool      `gorm:"not null;default:true" json:"skus"`
	Brands        bool      `gorm:"not null;default:true" json:"brands"`
	Categories    bool      `gorm:"not null;default:true" json:"categories"`
	SubCategories bool      `gorm:"not null;default:true" json:"sub_categories"`
	Materials     bool      `gorm:"not null;default:true" json:"materials"`
	Dimensions    bool      `gorm:"not null;default:true" json:"dimensions"`
	Images        bool      `gorm:"not null;default:true" json:"images"`
	Colors        bool      `gorm:"not null;default:true" json:"colors"`
	SKUOptions    bool      `gorm:"not null;default:true" json:"sku_options"`
	CreatedBy     string    `gorm:"size:120" json:"created_by"`
	UpdatedBy     string    `gorm:"size:120" json:"updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (m *MerchantFeatureConfig) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}
