package repository

import (
	"strings"

	"bsku001/backend/internal/entity"
	"bsku001/backend/internal/model"

	"gorm.io/gorm"
)

type SKUOptionRepository struct{ db *gorm.DB }

func NewSKUOptionRepository(db *gorm.DB) *SKUOptionRepository { return &SKUOptionRepository{db: db} }

func (r *SKUOptionRepository) ListGroups(merchantID string, query model.SKUOptionGroupListQuery) ([]entity.SKUOptionGroup, int64, int, int, error) {
	page, limit := normalizePage(query.Page, query.Limit)
	db := r.db.Model(&entity.SKUOptionGroup{}).Where("merchant_id = ?", merchantID)
	if query.SKUID != "" {
		db = db.Where("sku_id = ?", query.SKUID)
	}
	if query.Search != "" {
		like := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where("LOWER(code) LIKE ? OR LOWER(name_th) LIKE ? OR LOWER(name_en) LIKE ?", like, like, like)
	}
	if query.Status != "" {
		db = db.Where("is_active = ?", strings.EqualFold(query.Status, "active") || strings.EqualFold(query.Status, "true"))
	}
	var total int64
	db.Count(&total)
	var items []entity.SKUOptionGroup
	err := db.Preload("Options", func(tx *gorm.DB) *gorm.DB { return tx.Order("sort_order ASC, created_at ASC") }).Order("sort_order ASC, created_at ASC").Limit(limit).Offset((page - 1) * limit).Find(&items).Error
	return items, total, page, limit, err
}

func (r *SKUOptionRepository) ResolveSKUID(merchantID, idOrCode string) (string, error) {
	var item entity.SKU
	err := r.db.Select("id").Where("merchant_id = ? AND (id::text = ? OR sku_code = ?)", merchantID, idOrCode, idOrCode).First(&item).Error
	return item.ID.String(), err
}

func (r *SKUOptionRepository) ListGroupsBySKU(merchantID, skuID string) ([]entity.SKUOptionGroup, error) {
	var items []entity.SKUOptionGroup
	err := r.db.Where("merchant_id = ? AND sku_id = ? AND is_active = ?", merchantID, skuID, true).Preload("Options", func(tx *gorm.DB) *gorm.DB {
		return tx.Where("is_active = ?", true).Order("sort_order ASC, created_at ASC")
	}).Order("sort_order ASC, created_at ASC").Find(&items).Error
	return items, err
}

func (r *SKUOptionRepository) GetGroup(merchantID, id string) (entity.SKUOptionGroup, error) {
	var item entity.SKUOptionGroup
	err := r.db.Where("merchant_id = ? AND id::text = ?", merchantID, id).Preload("Options", func(tx *gorm.DB) *gorm.DB { return tx.Order("sort_order ASC, created_at ASC") }).First(&item).Error
	return item, err
}

func (r *SKUOptionRepository) CreateGroup(item entity.SKUOptionGroup, options []entity.SKUOptionValue) (entity.SKUOptionGroup, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&item).Error; err != nil {
			return err
		}
		for i := range options {
			options[i].OptionGroupID = item.ID.String()
			options[i].MerchantID = item.MerchantID
			options[i].CreatedBy = item.CreatedBy
			options[i].UpdatedBy = item.UpdatedBy
		}
		if len(options) > 0 {
			if err := tx.Create(&options).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return entity.SKUOptionGroup{}, err
	}
	return r.GetGroup(item.MerchantID, item.ID.String())
}

func (r *SKUOptionRepository) UpdateGroup(item entity.SKUOptionGroup, replaceOptions bool, options []entity.SKUOptionValue) (entity.SKUOptionGroup, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&item).Error; err != nil {
			return err
		}
		if replaceOptions {
			if err := tx.Where("merchant_id = ? AND option_group_id = ?", item.MerchantID, item.ID.String()).Delete(&entity.SKUOptionValue{}).Error; err != nil {
				return err
			}
			for i := range options {
				options[i].OptionGroupID = item.ID.String()
				options[i].MerchantID = item.MerchantID
				options[i].CreatedBy = item.UpdatedBy
				options[i].UpdatedBy = item.UpdatedBy
			}
			if len(options) > 0 {
				if err := tx.Create(&options).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return entity.SKUOptionGroup{}, err
	}
	return r.GetGroup(item.MerchantID, item.ID.String())
}

func (r *SKUOptionRepository) DeleteGroup(merchantID, id string) error {
	return r.db.Where("merchant_id = ? AND id::text = ?", merchantID, id).Delete(&entity.SKUOptionGroup{}).Error
}

func (r *SKUOptionRepository) CreateValue(item entity.SKUOptionValue) (entity.SKUOptionValue, error) {
	return item, r.db.Create(&item).Error
}

func (r *SKUOptionRepository) GetValue(merchantID, id string) (entity.SKUOptionValue, error) {
	var item entity.SKUOptionValue
	err := r.db.Where("merchant_id = ? AND id::text = ?", merchantID, id).First(&item).Error
	return item, err
}

func (r *SKUOptionRepository) UpdateValue(item entity.SKUOptionValue) (entity.SKUOptionValue, error) {
	return item, r.db.Save(&item).Error
}

func (r *SKUOptionRepository) DeleteValue(merchantID, id string) error {
	return r.db.Where("merchant_id = ? AND id::text = ?", merchantID, id).Delete(&entity.SKUOptionValue{}).Error
}
