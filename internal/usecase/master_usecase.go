package usecase

import (
	"math"

	"bsku001/backend/internal/entity"
	"bsku001/backend/internal/model"
	"bsku001/backend/internal/repository"
)

type MasterUsecase struct{ repo *repository.MasterRepository }

func NewMasterUsecase(repo *repository.MasterRepository) *MasterUsecase {
	return &MasterUsecase{repo: repo}
}

func (u *MasterUsecase) List(resource string, auth model.AuthContext, query model.ListQuery) (model.PageResponse[any], error) {
	items, total, err := u.repo.List(resource, auth.MerchantID, query)
	if err != nil {
		return model.PageResponse[any]{}, err
	}
	page, limit := normalizePage(query.Page, query.Limit)
	return model.PageResponse[any]{Items: mapMasterSlice(items), Page: page, Limit: limit, Total: total, TotalPages: int(math.Ceil(float64(total) / float64(limit)))}, nil
}

func (u *MasterUsecase) Get(resource string, auth model.AuthContext, id string) (any, error) {
	item, err := u.repo.Get(resource, auth.MerchantID, id)
	if err != nil {
		return nil, err
	}
	return mapMasterItem(item), nil
}

func (u *MasterUsecase) Create(resource string, auth model.AuthContext, req any) (any, error) {
	item, err := u.repo.Create(resource, auth.MerchantID, auth.UserID, req)
	if err != nil {
		return nil, err
	}
	return mapMasterItem(item), nil
}

func (u *MasterUsecase) Update(resource string, auth model.AuthContext, id string, req any) (any, error) {
	item, err := u.repo.Update(resource, auth.MerchantID, auth.UserID, id, req)
	if err != nil {
		return nil, err
	}
	return mapMasterItem(item), nil
}

func (u *MasterUsecase) Delete(resource string, auth model.AuthContext, id string) error {
	return u.repo.Delete(resource, auth.MerchantID, id)
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

func mapMasterSlice(value any) []any {
	out := []any{}
	switch items := value.(type) {
	case []entity.Brand:
		for _, item := range items {
			out = append(out, mapMasterItem(item))
		}
	case []entity.Category:
		for _, item := range items {
			out = append(out, mapMasterItem(item))
		}
	case []entity.SubCategory:
		for _, item := range items {
			out = append(out, mapMasterItem(item))
		}
	case []entity.Material:
		for _, item := range items {
			out = append(out, mapMasterItem(item))
		}
	case []entity.Color:
		for _, item := range items {
			out = append(out, mapMasterItem(item))
		}
	}
	return out
}

func mapMasterItem(value any) any {
	switch item := value.(type) {
	case entity.Brand:
		return model.MasterResponse{ID: item.ID.String(), MerchantID: item.MerchantID, Code: item.Code, NameTH: item.NameTH, NameEN: item.NameEN, IsActive: item.IsActive, AuditFields: model.AuditFields{CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt, CreatedBy: item.CreatedBy, UpdatedBy: item.UpdatedBy}}
	case entity.Category:
		return model.MasterResponse{ID: item.ID.String(), MerchantID: item.MerchantID, Code: item.Code, NameTH: item.NameTH, NameEN: item.NameEN, IsActive: item.IsActive, AuditFields: model.AuditFields{CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt, CreatedBy: item.CreatedBy, UpdatedBy: item.UpdatedBy}}
	case entity.SubCategory:
		return model.SubCategoryResponse{ID: item.ID.String(), MerchantID: item.MerchantID, CategoryID: item.CategoryID, Code: item.Code, NameTH: item.NameTH, NameEN: item.NameEN, IsActive: item.IsActive, AuditFields: model.AuditFields{CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt, CreatedBy: item.CreatedBy, UpdatedBy: item.UpdatedBy}}
	case entity.Material:
		return model.MasterResponse{ID: item.ID.String(), MerchantID: item.MerchantID, Code: item.Code, NameTH: item.NameTH, NameEN: item.NameEN, IsActive: item.IsActive, AuditFields: model.AuditFields{CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt, CreatedBy: item.CreatedBy, UpdatedBy: item.UpdatedBy}}
	case entity.Color:
		return model.MasterResponse{ID: item.ID.String(), MerchantID: item.MerchantID, Code: item.Code, NameTH: item.NameTH, NameEN: item.NameEN, HexCode: item.HexCode, IsActive: item.IsActive, AuditFields: model.AuditFields{CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt, CreatedBy: item.CreatedBy, UpdatedBy: item.UpdatedBy}}
	default:
		return value
	}
}
