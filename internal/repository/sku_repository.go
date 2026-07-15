package repository

import (
	"strings"

	"bsku001/backend/internal/entity"
	"bsku001/backend/internal/model"

	"gorm.io/gorm"
)

type SKURepository struct{ db *gorm.DB }

func NewSKURepository(db *gorm.DB) *SKURepository { return &SKURepository{db: db} }

func (r *SKURepository) List(merchantID string, query model.SKUListQuery) ([]entity.SKU, int64, int, int, error) {
	page, limit := normalizePage(query.Page, query.Limit)
	db := r.db.Model(&entity.SKU{}).Where("merchant_id = ?", merchantID)
	if query.Search != "" {
		like := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where("LOWER(sku_code) LIKE ? OR LOWER(gtin) LIKE ? OR LOWER(names::text) LIKE ? OR LOWER(descriptions::text) LIKE ?", like, like, like, like)
	}
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	if query.BrandID != "" {
		db = db.Where("brand_id = ?", query.BrandID)
	}
	if query.CategoryID != "" {
		db = db.Where("category_id = ?", query.CategoryID)
	}
	var total int64
	db.Count(&total)
	var items []entity.SKU
	err := db.Preload("Dimension").Preload("Images", func(tx *gorm.DB) *gorm.DB { return tx.Order("sort_order ASC, created_at ASC") }).Order("created_at DESC").Limit(limit).Offset((page - 1) * limit).Find(&items).Error
	return items, total, page, limit, err
}

func (r *SKURepository) Get(merchantID, idOrCode string) (entity.SKU, error) {
	var item entity.SKU
	err := r.db.Where("merchant_id = ? AND (id::text = ? OR sku_code = ?)", merchantID, idOrCode, idOrCode).Preload("Dimension").Preload("Images", func(tx *gorm.DB) *gorm.DB { return tx.Order("sort_order ASC, created_at ASC") }).First(&item).Error
	return item, err
}

func (r *SKURepository) Create(item entity.SKU, dimension entity.SKUDimension, images []entity.SKUImage) (entity.SKU, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&item).Error; err != nil {
			return err
		}
		dimension.SKUID = item.ID.String()
		dimension.MerchantID = item.MerchantID
		dimension.CreatedBy = item.CreatedBy
		dimension.UpdatedBy = item.UpdatedBy
		if err := tx.Create(&dimension).Error; err != nil {
			return err
		}
		for i := range images {
			images[i].SKUID = item.ID.String()
			images[i].MerchantID = item.MerchantID
			images[i].CreatedBy = item.CreatedBy
			images[i].UpdatedBy = item.UpdatedBy
		}
		if len(images) > 0 {
			if err := tx.Create(&images).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return entity.SKU{}, err
	}
	return r.Get(item.MerchantID, item.ID.String())
}

func (r *SKURepository) Update(item entity.SKU, dimension entity.SKUDimension, replaceImages bool, images []entity.SKUImage) (entity.SKU, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&item).Error; err != nil {
			return err
		}
		var existing entity.SKUDimension
		if err := tx.Where("merchant_id = ? AND sku_id = ?", item.MerchantID, item.ID.String()).First(&existing).Error; err == nil {
			dimension.ID = existing.ID
			dimension.MerchantID = existing.MerchantID
			dimension.SKUID = existing.SKUID
			dimension.CreatedBy = existing.CreatedBy
			dimension.UpdatedBy = item.UpdatedBy
			if err := tx.Save(&dimension).Error; err != nil {
				return err
			}
		} else {
			dimension.SKUID = item.ID.String()
			dimension.MerchantID = item.MerchantID
			dimension.CreatedBy = item.CreatedBy
			dimension.UpdatedBy = item.UpdatedBy
			if err := tx.Create(&dimension).Error; err != nil {
				return err
			}
		}
		if replaceImages {
			if err := tx.Where("merchant_id = ? AND sku_id = ?", item.MerchantID, item.ID.String()).Delete(&entity.SKUImage{}).Error; err != nil {
				return err
			}
			for i := range images {
				images[i].SKUID = item.ID.String()
				images[i].MerchantID = item.MerchantID
				images[i].CreatedBy = item.UpdatedBy
				images[i].UpdatedBy = item.UpdatedBy
			}
			if len(images) > 0 {
				if err := tx.Create(&images).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return entity.SKU{}, err
	}
	return r.Get(item.MerchantID, item.ID.String())
}

func (r *SKURepository) SoftDelete(merchantID, id string) error {
	return r.db.Model(&entity.SKU{}).Where("merchant_id = ? AND id::text = ?", merchantID, id).Updates(map[string]any{"is_active": false, "status": "DISCONTINUED"}).Error
}

func (r *SKURepository) GetDimension(merchantID, skuID string) (entity.SKUDimension, error) {
	var item entity.SKUDimension
	err := r.db.Where("merchant_id = ? AND sku_id = ?", merchantID, skuID).First(&item).Error
	return item, err
}

func (r *SKURepository) UpsertDimension(item entity.SKUDimension) (entity.SKUDimension, error) {
	var existing entity.SKUDimension
	if err := r.db.Where("merchant_id = ? AND sku_id = ?", item.MerchantID, item.SKUID).First(&existing).Error; err == nil {
		item.ID = existing.ID
		item.CreatedBy = existing.CreatedBy
		if err := r.db.Save(&item).Error; err != nil {
			return entity.SKUDimension{}, err
		}
		return item, nil
	}
	if err := r.db.Create(&item).Error; err != nil {
		return entity.SKUDimension{}, err
	}
	return item, nil
}

func (r *SKURepository) ListImages(merchantID, skuID string) ([]entity.SKUImage, error) {
	var items []entity.SKUImage
	err := r.db.Where("merchant_id = ? AND sku_id = ?", merchantID, skuID).Order("sort_order ASC, created_at ASC").Find(&items).Error
	return items, err
}

func (r *SKURepository) CreateImage(item entity.SKUImage) (entity.SKUImage, error) {
	return item, r.db.Create(&item).Error
}

func (r *SKURepository) UpdateImage(merchantID, imageID string, item entity.SKUImage) (entity.SKUImage, error) {
	var existing entity.SKUImage
	if err := r.db.Where("merchant_id = ? AND id::text = ?", merchantID, imageID).First(&existing).Error; err != nil {
		return entity.SKUImage{}, err
	}
	existing.ImageURL = item.ImageURL
	existing.AltText = item.AltText
	existing.SortOrder = item.SortOrder
	existing.IsPrimary = item.IsPrimary
	existing.UpdatedBy = item.UpdatedBy
	return existing, r.db.Save(&existing).Error
}

func (r *SKURepository) DeleteImage(merchantID, imageID string) error {
	return r.db.Where("merchant_id = ? AND id::text = ?", merchantID, imageID).Delete(&entity.SKUImage{}).Error
}
