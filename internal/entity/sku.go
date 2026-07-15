package entity

import "gorm.io/datatypes"

type SKU struct {
	Base
	SKUCode         string         `gorm:"size:120;not null;uniqueIndex:idx_sku_merchant_code" json:"sku_code"`
	GTIN            string         `gorm:"size:80;index" json:"gtin"`
	Names           datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'" json:"names"`
	Descriptions    datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'" json:"descriptions"`
	BrandID         string         `gorm:"size:80;index" json:"brand_id"`
	CategoryID      string         `gorm:"type:uuid;index" json:"category_id"`
	SubCategoryID   string         `gorm:"size:80;index" json:"sub_category_id"`
	MaterialID      string         `gorm:"type:uuid;index" json:"material_id"`
	ColorID         string         `gorm:"size:80;index" json:"color_id"`
	Status          string         `gorm:"size:40;not null;default:'DRAFT';index" json:"status"`
	IsHazmat        bool           `gorm:"not null;default:false" json:"is_hazmat"`
	CountryOfOrigin string         `gorm:"size:2" json:"country_of_origin"`
	IsActive        bool           `gorm:"not null;default:true;index" json:"is_active"`
	Dimension       SKUDimension   `gorm:"foreignKey:SKUID" json:"dimension,omitempty"`
	Images          []SKUImage     `gorm:"foreignKey:SKUID" json:"images,omitempty"`
}

type SKUDimension struct {
	Base
	SKUID      string  `gorm:"type:uuid;not null;uniqueIndex:idx_dimension_sku" json:"sku_id"`
	Weight     float64 `gorm:"type:decimal(10,4);not null;default:0" json:"weight"`
	WeightUnit string  `gorm:"size:20;not null;default:'kg'" json:"weight_unit"`
	Width      float64 `gorm:"type:decimal(10,4);not null;default:0" json:"width"`
	Length     float64 `gorm:"type:decimal(10,4);not null;default:0" json:"length"`
	Height     float64 `gorm:"type:decimal(10,4);not null;default:0" json:"height"`
	Unit       string  `gorm:"size:20;not null;default:'cm'" json:"unit"`
	HSCode     string  `gorm:"size:80" json:"hs_code"`
}

type SKUImage struct {
	Base
	SKUID     string `gorm:"type:uuid;not null;index" json:"sku_id"`
	ImageURL  string `gorm:"size:1000;not null" json:"image_url"`
	AltText   string `gorm:"size:255" json:"alt_text"`
	SortOrder int    `gorm:"not null;default:0" json:"sort_order"`
	IsPrimary bool   `gorm:"not null;default:false" json:"is_primary"`
}
