package entity

type SKUOptionGroup struct {
	Base
	SKUID         string           `gorm:"type:uuid;not null;index:idx_sku_option_groups_sku" json:"sku_id"`
	Code          string           `gorm:"size:80;not null;index" json:"code"`
	NameTH        string           `gorm:"size:255;not null" json:"name_th"`
	NameEN        string           `gorm:"size:255;not null" json:"name_en"`
	SelectionType string           `gorm:"size:20;not null;default:'RADIO';index" json:"selection_type"`
	IsRequired    bool             `gorm:"not null;default:false" json:"is_required"`
	MinSelect     int              `gorm:"not null;default:0" json:"min_select"`
	MaxSelect     int              `gorm:"not null;default:1" json:"max_select"`
	SortOrder     int              `gorm:"not null;default:0" json:"sort_order"`
	IsActive      bool             `gorm:"not null;default:true;index" json:"is_active"`
	Options       []SKUOptionValue `gorm:"foreignKey:OptionGroupID" json:"options,omitempty"`
}

type SKUOptionValue struct {
	Base
	OptionGroupID string  `gorm:"type:uuid;not null;index:idx_sku_option_values_group" json:"option_group_id"`
	Code          string  `gorm:"size:80;not null;index" json:"code"`
	NameTH        string  `gorm:"size:255;not null" json:"name_th"`
	NameEN        string  `gorm:"size:255;not null" json:"name_en"`
	PriceDelta    float64 `gorm:"type:decimal(10,4);not null;default:0" json:"price_delta"`
	SortOrder     int     `gorm:"not null;default:0" json:"sort_order"`
	IsDefault     bool    `gorm:"not null;default:false" json:"is_default"`
	IsActive      bool    `gorm:"not null;default:true;index" json:"is_active"`
}
