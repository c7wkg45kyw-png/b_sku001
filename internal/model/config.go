package model

type MerchantFeatureConfigResponse struct {
	MerchantID    string `json:"merchant_id"`
	SKUs          bool   `json:"skus"`
	Brands        bool   `json:"brands"`
	Categories    bool   `json:"categories"`
	SubCategories bool   `json:"sub_categories"`
	Materials     bool   `json:"materials"`
	Dimensions    bool   `json:"dimensions"`
	Images        bool   `json:"images"`
	Colors        bool   `json:"colors"`
	SKUOptions    bool   `json:"sku_options"`
}

type MerchantFeatureConfigUpdateRequest struct {
	SKUs          *bool `json:"skus,omitempty"`
	Brands        *bool `json:"brands,omitempty"`
	Categories    *bool `json:"categories,omitempty"`
	SubCategories *bool `json:"sub_categories,omitempty"`
	Materials     *bool `json:"materials,omitempty"`
	Dimensions    *bool `json:"dimensions,omitempty"`
	Images        *bool `json:"images,omitempty"`
	Colors        *bool `json:"colors,omitempty"`
	SKUOptions    *bool `json:"sku_options,omitempty"`
}
