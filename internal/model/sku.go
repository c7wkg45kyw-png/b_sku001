package model

type LocalizedText map[string]string

type WeightPayload struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type DimensionPayload struct {
	Width  UnitValue `json:"width"`
	Length UnitValue `json:"length"`
	Height UnitValue `json:"height"`
	HSCode string    `json:"hs_code"`
}

type UnitValue struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type SKUDimensionRequest struct {
	Weight     float64 `json:"weight"`
	WeightUnit string  `json:"weight_unit"`
	Width      float64 `json:"width"`
	Length     float64 `json:"length"`
	Height     float64 `json:"height"`
	Unit       string  `json:"unit"`
	HSCode     string  `json:"hs_code"`
}

type SKUDimensionResponse struct {
	ID         string  `json:"id"`
	SKUID      string  `json:"sku_id"`
	Weight     float64 `json:"weight"`
	WeightUnit string  `json:"weight_unit"`
	Width      float64 `json:"width"`
	Length     float64 `json:"length"`
	Height     float64 `json:"height"`
	Unit       string  `json:"unit"`
	HSCode     string  `json:"hs_code"`
	AuditFields
}

type SKUImageRequest struct {
	ImageURL  string `json:"image_url" binding:"required"`
	AltText   string `json:"alt_text"`
	SortOrder int    `json:"sort_order"`
	IsPrimary bool   `json:"is_primary"`
}

type SKURequest struct {
	SKUCode         string            `json:"sku_code" binding:"required"`
	GTIN            string            `json:"gtin"`
	BrandID         string            `json:"brand_id"`
	CategoryID      string            `json:"category_id"`
	SubCategoryID   string            `json:"sub_category_id"`
	MaterialID      string            `json:"material_id"`
	ColorID         string            `json:"color_id"`
	Names           LocalizedText     `json:"names" binding:"required"`
	Descriptions    LocalizedText     `json:"descriptions"`
	Weight          WeightPayload     `json:"weight"`
	Dimensions      DimensionPayload  `json:"dimensions"`
	Status          string            `json:"status"`
	IsHazmat        bool              `json:"is_hazmat"`
	CountryOfOrigin string            `json:"country_of_origin"`
	Images          []SKUImageRequest `json:"images"`
}

type SKUPatchRequest struct {
	SKUCode         *string           `json:"sku_code"`
	GTIN            *string           `json:"gtin"`
	BrandID         *string           `json:"brand_id"`
	CategoryID      *string           `json:"category_id"`
	SubCategoryID   *string           `json:"sub_category_id"`
	MaterialID      *string           `json:"material_id"`
	ColorID         *string           `json:"color_id"`
	Names           LocalizedText     `json:"names"`
	Descriptions    LocalizedText     `json:"descriptions"`
	Weight          *WeightPayload    `json:"weight"`
	Dimensions      *DimensionPayload `json:"dimensions"`
	Status          *string           `json:"status"`
	IsHazmat        *bool             `json:"is_hazmat"`
	CountryOfOrigin *string           `json:"country_of_origin"`
	IsActive        *bool             `json:"is_active"`
}

type SKUResponse struct {
	ID              string             `json:"id"`
	MerchantID      string             `json:"merchant_id"`
	SKUCode         string             `json:"sku_code"`
	GTIN            string             `json:"gtin"`
	BrandID         string             `json:"brand_id"`
	CategoryID      string             `json:"category_id"`
	SubCategoryID   string             `json:"sub_category_id"`
	MaterialID      string             `json:"material_id"`
	ColorID         string             `json:"color_id"`
	Names           LocalizedText      `json:"names,omitempty"`
	Name            string             `json:"name,omitempty"`
	Descriptions    LocalizedText      `json:"descriptions,omitempty"`
	Description     string             `json:"description,omitempty"`
	Weight          WeightPayload      `json:"weight"`
	Dimensions      DimensionPayload   `json:"dimensions"`
	HSCode          string             `json:"hs_code"`
	Status          string             `json:"status"`
	IsHazmat        bool               `json:"is_hazmat"`
	CountryOfOrigin string             `json:"country_of_origin"`
	IsActive        bool               `json:"is_active"`
	Images          []SKUImageResponse `json:"images"`
	AuditFields
}

type SKUImageResponse struct {
	ID        string `json:"id"`
	ImageURL  string `json:"image_url"`
	AltText   string `json:"alt_text"`
	SortOrder int    `json:"sort_order"`
	IsPrimary bool   `json:"is_primary"`
}

type SKUListQuery struct {
	ListQuery
	BrandID    string `form:"brand_id"`
	CategoryID string `form:"category_id"`
}
