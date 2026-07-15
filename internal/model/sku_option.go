package model

type SKUOptionValueRequest struct {
	Code       string  `json:"code" binding:"required"`
	NameTH     string  `json:"name_th" binding:"required"`
	NameEN     string  `json:"name_en" binding:"required"`
	PriceDelta float64 `json:"price_delta"`
	SortOrder  int     `json:"sort_order"`
	IsDefault  bool    `json:"is_default"`
}

type SKUOptionGroupRequest struct {
	Code          string                  `json:"code" binding:"required"`
	NameTH        string                  `json:"name_th" binding:"required"`
	NameEN        string                  `json:"name_en" binding:"required"`
	SelectionType string                  `json:"selection_type" binding:"required"`
	IsRequired    bool                    `json:"is_required"`
	MinSelect     int                     `json:"min_select"`
	MaxSelect     int                     `json:"max_select"`
	SortOrder     int                     `json:"sort_order"`
	Options       []SKUOptionValueRequest `json:"options"`
}

type SKUOptionValueResponse struct {
	ID            string  `json:"id"`
	MerchantID    string  `json:"merchant_id"`
	OptionGroupID string  `json:"option_group_id"`
	Code          string  `json:"code"`
	NameTH        string  `json:"name_th"`
	NameEN        string  `json:"name_en"`
	Name          string  `json:"name,omitempty"`
	PriceDelta    float64 `json:"price_delta"`
	SortOrder     int     `json:"sort_order"`
	IsDefault     bool    `json:"is_default"`
	IsActive      bool    `json:"is_active"`
	AuditFields
}

type SKUOptionGroupResponse struct {
	ID            string                   `json:"id"`
	MerchantID    string                   `json:"merchant_id"`
	SKUID         string                   `json:"sku_id"`
	Code          string                   `json:"code"`
	NameTH        string                   `json:"name_th"`
	NameEN        string                   `json:"name_en"`
	Name          string                   `json:"name,omitempty"`
	SelectionType string                   `json:"selection_type"`
	IsRequired    bool                     `json:"is_required"`
	MinSelect     int                      `json:"min_select"`
	MaxSelect     int                      `json:"max_select"`
	SortOrder     int                      `json:"sort_order"`
	IsActive      bool                     `json:"is_active"`
	Options       []SKUOptionValueResponse `json:"options"`
	AuditFields
}

type SKUOptionGroupListQuery struct {
	ListQuery
	SKUID string `form:"sku_id"`
}
