package usecase

import (
	"bsku001/backend/internal/entity"
	"bsku001/backend/internal/model"
)

type SKURepositoryPort interface {
	List(merchantID string, query model.SKUListQuery) ([]entity.SKU, int64, int, int, error)
	Get(merchantID, idOrCode string) (entity.SKU, error)
	Create(item entity.SKU, dimension entity.SKUDimension, images []entity.SKUImage) (entity.SKU, error)
	Update(item entity.SKU, dimension entity.SKUDimension, replaceImages bool, images []entity.SKUImage) (entity.SKU, error)
	SoftDelete(merchantID, id string) error
	GetDimension(merchantID, skuID string) (entity.SKUDimension, error)
	UpsertDimension(item entity.SKUDimension) (entity.SKUDimension, error)
	ListImages(merchantID, skuID string) ([]entity.SKUImage, error)
	CreateImage(item entity.SKUImage) (entity.SKUImage, error)
	UpdateImage(merchantID, imageID string, item entity.SKUImage) (entity.SKUImage, error)
	DeleteImage(merchantID, imageID string) error
}

type MasterRepositoryPort interface {
	List(resource, merchantID string, query model.ListQuery) (any, int64, error)
	Get(resource, merchantID, id string) (any, error)
	Create(resource, merchantID, userID string, req any) (any, error)
	Update(resource, merchantID, userID, id string, req any) (any, error)
	Delete(resource, merchantID, id string) error
}
