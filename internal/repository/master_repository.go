package repository

import (
	"errors"
	"strings"

	"bsku001/backend/internal/entity"
	"bsku001/backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterRepository struct {
	db *gorm.DB
}

func NewMasterRepository(db *gorm.DB) *MasterRepository { return &MasterRepository{db: db} }

func (r *MasterRepository) List(resource, merchantID string, query model.ListQuery) (any, int64, error) {
	page, limit := normalizePage(query.Page, query.Limit)
	offset := (page - 1) * limit
	db := r.db.Where("merchant_id = ?", merchantID)
	if query.Search != "" {
		like := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where("LOWER(code) LIKE ? OR LOWER(name_th) LIKE ? OR LOWER(name_en) LIKE ?", like, like, like)
	}
	switch resource {
	case "brands":
		var items []entity.Brand
		var total int64
		db.Model(&entity.Brand{}).Count(&total)
		err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&items).Error
		return items, total, err
	case "categories":
		var items []entity.Category
		var total int64
		db.Model(&entity.Category{}).Count(&total)
		err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&items).Error
		return items, total, err
	case "sub-categories":
		var items []entity.SubCategory
		var total int64
		db.Model(&entity.SubCategory{}).Count(&total)
		err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&items).Error
		return items, total, err
	case "materials":
		var items []entity.Material
		var total int64
		db.Model(&entity.Material{}).Count(&total)
		err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&items).Error
		return items, total, err
	case "colors":
		var items []entity.Color
		var total int64
		db.Model(&entity.Color{}).Count(&total)
		err := db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&items).Error
		return items, total, err
	default:
		return nil, 0, errors.New("unknown master resource")
	}
}

func (r *MasterRepository) Get(resource, merchantID, id string) (any, error) {
	db := r.db.Where("merchant_id = ? AND id = ?", merchantID, id)
	switch resource {
	case "brands":
		var item entity.Brand
		return item, db.First(&item).Error
	case "categories":
		var item entity.Category
		return item, db.First(&item).Error
	case "sub-categories":
		var item entity.SubCategory
		return item, db.First(&item).Error
	case "materials":
		var item entity.Material
		return item, db.First(&item).Error
	case "colors":
		var item entity.Color
		return item, db.First(&item).Error
	default:
		return nil, errors.New("unknown master resource")
	}
}

func (r *MasterRepository) Create(resource, merchantID, userID string, req any) (any, error) {
	switch resource {
	case "brands":
		payload := req.(model.MasterRequest)
		item := entity.Brand{Base: entity.Base{MerchantID: merchantID, CreatedBy: userID, UpdatedBy: userID}, Code: payload.Code, NameTH: payload.NameTH, NameEN: payload.NameEN, IsActive: true}
		return item, r.db.Create(&item).Error
	case "categories":
		payload := req.(model.MasterRequest)
		item := entity.Category{Base: entity.Base{MerchantID: merchantID, CreatedBy: userID, UpdatedBy: userID}, Code: payload.Code, NameTH: payload.NameTH, NameEN: payload.NameEN, IsActive: true}
		return item, r.db.Create(&item).Error
	case "sub-categories":
		payload := req.(model.SubCategoryRequest)
		item := entity.SubCategory{Base: entity.Base{MerchantID: merchantID, CreatedBy: userID, UpdatedBy: userID}, CategoryID: payload.CategoryID, Code: payload.Code, NameTH: payload.NameTH, NameEN: payload.NameEN, IsActive: true}
		return item, r.db.Create(&item).Error
	case "materials":
		payload := req.(model.MasterRequest)
		item := entity.Material{Base: entity.Base{MerchantID: merchantID, CreatedBy: userID, UpdatedBy: userID}, Code: payload.Code, NameTH: payload.NameTH, NameEN: payload.NameEN, IsActive: true}
		return item, r.db.Create(&item).Error
	case "colors":
		payload := req.(model.MasterRequest)
		item := entity.Color{Base: entity.Base{MerchantID: merchantID, CreatedBy: userID, UpdatedBy: userID}, Code: payload.Code, NameTH: payload.NameTH, NameEN: payload.NameEN, HexCode: payload.HexCode, IsActive: true}
		return item, r.db.Create(&item).Error
	default:
		return nil, errors.New("unknown master resource")
	}
}

func (r *MasterRepository) Update(resource, merchantID, userID, id string, req any) (any, error) {
	item, err := r.Get(resource, merchantID, id)
	if err != nil {
		return nil, err
	}
	switch current := item.(type) {
	case entity.Brand:
		payload := req.(model.MasterRequest)
		current.Code = payload.Code
		current.NameTH = payload.NameTH
		current.NameEN = payload.NameEN
		current.UpdatedBy = userID
		err = r.db.Save(&current).Error
		return current, err
	case entity.Category:
		payload := req.(model.MasterRequest)
		current.Code = payload.Code
		current.NameTH = payload.NameTH
		current.NameEN = payload.NameEN
		current.UpdatedBy = userID
		err = r.db.Save(&current).Error
		return current, err
	case entity.SubCategory:
		payload := req.(model.SubCategoryRequest)
		current.CategoryID = payload.CategoryID
		current.Code = payload.Code
		current.NameTH = payload.NameTH
		current.NameEN = payload.NameEN
		current.UpdatedBy = userID
		err = r.db.Save(&current).Error
		return current, err
	case entity.Material:
		payload := req.(model.MasterRequest)
		current.Code = payload.Code
		current.NameTH = payload.NameTH
		current.NameEN = payload.NameEN
		current.UpdatedBy = userID
		err = r.db.Save(&current).Error
		return current, err
	case entity.Color:
		payload := req.(model.MasterRequest)
		current.Code = payload.Code
		current.NameTH = payload.NameTH
		current.NameEN = payload.NameEN
		current.HexCode = payload.HexCode
		current.UpdatedBy = userID
		err = r.db.Save(&current).Error
		return current, err
	default:
		return nil, errors.New("unknown master resource")
	}
}

func (r *MasterRepository) Delete(resource, merchantID, id string) error {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	switch resource {
	case "brands":
		return r.db.Where("merchant_id = ?", merchantID).Delete(&entity.Brand{}, parsed).Error
	case "categories":
		return r.db.Where("merchant_id = ?", merchantID).Delete(&entity.Category{}, parsed).Error
	case "sub-categories":
		return r.db.Where("merchant_id = ?", merchantID).Delete(&entity.SubCategory{}, parsed).Error
	case "materials":
		return r.db.Where("merchant_id = ?", merchantID).Delete(&entity.Material{}, parsed).Error
	case "colors":
		return r.db.Where("merchant_id = ?", merchantID).Delete(&entity.Color{}, parsed).Error
	default:
		return errors.New("unknown master resource")
	}
}

func normalizePage(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	return page, limit
}
