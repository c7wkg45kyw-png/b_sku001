package model

type MasterRequest struct {
	Code     string `json:"code" binding:"required"`
	NameTH   string `json:"name_th" binding:"required"`
	NameEN   string `json:"name_en" binding:"required"`
	HexCode  string `json:"hex_code,omitempty"`
	IsActive *bool  `json:"is_active,omitempty"`
}

type SubCategoryRequest struct {
	CategoryID string `json:"category_id" binding:"required"`
	Code       string `json:"code" binding:"required"`
	NameTH     string `json:"name_th" binding:"required"`
	NameEN     string `json:"name_en" binding:"required"`
	IsActive   *bool  `json:"is_active,omitempty"`
}

type MasterResponse struct {
	ID         string `json:"id"`
	MerchantID string `json:"merchant_id"`
	Code       string `json:"code"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	HexCode    string `json:"hex_code,omitempty"`
	IsActive   bool   `json:"is_active"`
	AuditFields
}

type SubCategoryResponse struct {
	ID         string `json:"id"`
	MerchantID string `json:"merchant_id"`
	CategoryID string `json:"category_id"`
	Code       string `json:"code"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	IsActive   bool   `json:"is_active"`
	AuditFields
}
