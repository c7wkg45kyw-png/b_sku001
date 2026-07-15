package usecase

import (
	"bsku001/backend/internal/entity"
	"bsku001/backend/internal/model"
	"bsku001/backend/internal/repository"
)

type ConfigRepositoryPort interface {
	GetMerchantFeatureConfig(merchantID, userID string) (entity.MerchantFeatureConfig, error)
}

type ConfigUsecase struct {
	repo *repository.ConfigRepository
}

func NewConfigUsecase(repo *repository.ConfigRepository) *ConfigUsecase {
	return &ConfigUsecase{repo: repo}
}

func (u *ConfigUsecase) GetMerchantFeatureConfig(auth model.AuthContext) (model.MerchantFeatureConfigResponse, error) {
	item, err := u.repo.GetMerchantFeatureConfig(auth.MerchantID, auth.UserID)
	if err != nil {
		return model.MerchantFeatureConfigResponse{}, err
	}
	return model.MerchantFeatureConfigResponse{
		MerchantID:    item.MerchantID,
		SKUs:          item.SKUs,
		Brands:        item.Brands,
		Categories:    item.Categories,
		SubCategories: item.SubCategories,
		Materials:     item.Materials,
		Dimensions:    item.Dimensions,
		Images:        item.Images,
		Colors:        item.Colors,
		SKUOptions:    item.SKUOptions,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}, nil
}

func (u *ConfigUsecase) UpdateMerchantFeatureConfig(auth model.AuthContext, req model.MerchantFeatureConfigUpdateRequest) (model.MerchantFeatureConfigResponse, error) {
	item, err := u.repo.GetMerchantFeatureConfig(auth.MerchantID, auth.UserID)
	if err != nil {
		return model.MerchantFeatureConfigResponse{}, err
	}
	if req.SKUs != nil {
		item.SKUs = *req.SKUs
	}
	if req.Brands != nil {
		item.Brands = *req.Brands
	}
	if req.Categories != nil {
		item.Categories = *req.Categories
	}
	if req.SubCategories != nil {
		item.SubCategories = *req.SubCategories
	}
	if req.Materials != nil {
		item.Materials = *req.Materials
	}
	if req.Dimensions != nil {
		item.Dimensions = *req.Dimensions
	}
	if req.Images != nil {
		item.Images = *req.Images
	}
	if req.Colors != nil {
		item.Colors = *req.Colors
	}
	if req.SKUOptions != nil {
		item.SKUOptions = *req.SKUOptions
	}
	item.UpdatedBy = auth.UserID
	updated, err := u.repo.UpdateMerchantFeatureConfig(item)
	if err != nil {
		return model.MerchantFeatureConfigResponse{}, err
	}
	return model.MerchantFeatureConfigResponse{
		MerchantID:    updated.MerchantID,
		SKUs:          updated.SKUs,
		Brands:        updated.Brands,
		Categories:    updated.Categories,
		SubCategories: updated.SubCategories,
		Materials:     updated.Materials,
		Dimensions:    updated.Dimensions,
		Images:        updated.Images,
		Colors:        updated.Colors,
		SKUOptions:    updated.SKUOptions,
		CreatedAt:     updated.CreatedAt,
		UpdatedAt:     updated.UpdatedAt,
	}, nil
}
