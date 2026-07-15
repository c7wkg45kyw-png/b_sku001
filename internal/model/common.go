package model

import "time"

type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type PageResponse[T any] struct {
	Items      []T   `json:"items"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type ListQuery struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"search"`
	Status string `form:"status"`
}

type AuditFields struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	CreatedBy string    `json:"-"`
	UpdatedBy string    `json:"-"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SummaryResponse struct {
	SumSKUs          int64 `json:"sum_skus"`
	SumBrand         int64 `json:"sum_brand"`
	SumCategories    int64 `json:"sum_categories"`
	SumSubCategories int64 `json:"sum_sub_categories"`
	SumMaterials     int64 `json:"sum_materials"`
	SumColors        int64 `json:"sum_colors"`
	SumOptions       int64 `json:"sum_options"`
}
