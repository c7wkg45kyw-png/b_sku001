package entity

type Brand struct {
	Base
	Code     string `gorm:"size:80;not null;uniqueIndex:idx_brand_merchant_code" json:"code"`
	NameTH   string `gorm:"size:255;not null" json:"name_th"`
	NameEN   string `gorm:"size:255;not null" json:"name_en"`
	IsActive bool   `gorm:"not null;default:true" json:"is_active"`
}

type Category struct {
	Base
	Code     string `gorm:"size:80;not null;uniqueIndex:idx_category_merchant_code" json:"code"`
	NameTH   string `gorm:"size:255;not null" json:"name_th"`
	NameEN   string `gorm:"size:255;not null" json:"name_en"`
	IsActive bool   `gorm:"not null;default:true" json:"is_active"`
}

type SubCategory struct {
	Base
	CategoryID string `gorm:"type:uuid;not null;index" json:"category_id"`
	Code       string `gorm:"size:80;not null;uniqueIndex:idx_subcategory_merchant_code" json:"code"`
	NameTH     string `gorm:"size:255;not null" json:"name_th"`
	NameEN     string `gorm:"size:255;not null" json:"name_en"`
	IsActive   bool   `gorm:"not null;default:true" json:"is_active"`
}

type Material struct {
	Base
	Code     string `gorm:"size:80;not null;uniqueIndex:idx_material_merchant_code" json:"code"`
	NameTH   string `gorm:"size:255;not null" json:"name_th"`
	NameEN   string `gorm:"size:255;not null" json:"name_en"`
	IsActive bool   `gorm:"not null;default:true" json:"is_active"`
}

type Color struct {
	Base
	Code     string `gorm:"size:80;not null;uniqueIndex:idx_color_merchant_code" json:"code"`
	NameTH   string `gorm:"size:255;not null" json:"name_th"`
	NameEN   string `gorm:"size:255;not null" json:"name_en"`
	HexCode  string `gorm:"size:20" json:"hex_code"`
	IsActive bool   `gorm:"not null;default:true" json:"is_active"`
}
