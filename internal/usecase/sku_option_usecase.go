package usecase

import (
	"errors"
	"math"
	"strings"

	"bsku001/backend/internal/entity"
	"bsku001/backend/internal/model"
	"bsku001/backend/internal/repository"
)

const (
	OptionSelectionRadio    = "RADIO"
	OptionSelectionCheckbox = "CHECKBOX"
)

type SKUOptionUsecase struct {
	repo *repository.SKUOptionRepository
}

func NewSKUOptionUsecase(repo *repository.SKUOptionRepository) *SKUOptionUsecase {
	return &SKUOptionUsecase{repo: repo}
}

func (u *SKUOptionUsecase) ListGroups(auth model.AuthContext, query model.SKUOptionGroupListQuery, lang string) (model.PageResponse[model.SKUOptionGroupResponse], error) {
	items, total, page, limit, err := u.repo.ListGroups(auth.MerchantID, query)
	if err != nil {
		return model.PageResponse[model.SKUOptionGroupResponse]{}, err
	}
	out := make([]model.SKUOptionGroupResponse, 0, len(items))
	for _, item := range items {
		out = append(out, mapOptionGroup(item, lang))
	}
	return model.PageResponse[model.SKUOptionGroupResponse]{Items: out, Page: page, Limit: limit, Total: total, TotalPages: int(math.Ceil(float64(total) / float64(limit)))}, nil
}

func (u *SKUOptionUsecase) ListGroupsBySKU(auth model.AuthContext, skuID, lang string) ([]model.SKUOptionGroupResponse, error) {
	resolvedSKUID, err := u.repo.ResolveSKUID(auth.MerchantID, skuID)
	if err != nil {
		return nil, err
	}
	items, err := u.repo.ListGroupsBySKU(auth.MerchantID, resolvedSKUID)
	if err != nil {
		return nil, err
	}
	out := make([]model.SKUOptionGroupResponse, 0, len(items))
	for _, item := range items {
		out = append(out, mapOptionGroup(item, lang))
	}
	return out, nil
}

func (u *SKUOptionUsecase) GetGroup(auth model.AuthContext, id, lang string) (model.SKUOptionGroupResponse, error) {
	item, err := u.repo.GetGroup(auth.MerchantID, id)
	if err != nil {
		return model.SKUOptionGroupResponse{}, err
	}
	return mapOptionGroup(item, lang), nil
}

func (u *SKUOptionUsecase) CreateGroup(auth model.AuthContext, skuID string, req model.SKUOptionGroupRequest, lang string) (model.SKUOptionGroupResponse, error) {
	if err := validateOptionGroupRequest(&req); err != nil {
		return model.SKUOptionGroupResponse{}, err
	}
	resolvedSKUID, err := u.repo.ResolveSKUID(auth.MerchantID, skuID)
	if err != nil {
		return model.SKUOptionGroupResponse{}, err
	}
	item := optionGroupFromRequest(auth, resolvedSKUID, req)
	created, err := u.repo.CreateGroup(item, optionValuesFromRequest(req.Options))
	if err != nil {
		return model.SKUOptionGroupResponse{}, err
	}
	return mapOptionGroup(created, lang), nil
}

func (u *SKUOptionUsecase) ReplaceGroup(auth model.AuthContext, id string, req model.SKUOptionGroupRequest, lang string) (model.SKUOptionGroupResponse, error) {
	if err := validateOptionGroupRequest(&req); err != nil {
		return model.SKUOptionGroupResponse{}, err
	}
	item, err := u.repo.GetGroup(auth.MerchantID, id)
	if err != nil {
		return model.SKUOptionGroupResponse{}, err
	}
	applyOptionGroupRequest(&item, req)
	item.UpdatedBy = auth.UserID
	updated, err := u.repo.UpdateGroup(item, true, optionValuesFromRequest(req.Options))
	if err != nil {
		return model.SKUOptionGroupResponse{}, err
	}
	return mapOptionGroup(updated, lang), nil
}

func (u *SKUOptionUsecase) DeleteGroup(auth model.AuthContext, id string) error {
	return u.repo.DeleteGroup(auth.MerchantID, id)
}

func (u *SKUOptionUsecase) CreateValue(auth model.AuthContext, groupID string, req model.SKUOptionValueRequest, lang string) (model.SKUOptionValueResponse, error) {
	if _, err := u.repo.GetGroup(auth.MerchantID, groupID); err != nil {
		return model.SKUOptionValueResponse{}, err
	}
	item := optionValueFromRequest(req)
	item.Base = entity.Base{MerchantID: auth.MerchantID, CreatedBy: auth.UserID, UpdatedBy: auth.UserID}
	item.OptionGroupID = groupID
	created, err := u.repo.CreateValue(item)
	if err != nil {
		return model.SKUOptionValueResponse{}, err
	}
	return mapOptionValue(created, lang), nil
}

func (u *SKUOptionUsecase) ReplaceValue(auth model.AuthContext, id string, req model.SKUOptionValueRequest, lang string) (model.SKUOptionValueResponse, error) {
	item, err := u.repo.GetValue(auth.MerchantID, id)
	if err != nil {
		return model.SKUOptionValueResponse{}, err
	}
	updated := optionValueFromRequest(req)
	updated.ID = item.ID
	updated.MerchantID = item.MerchantID
	updated.OptionGroupID = item.OptionGroupID
	updated.CreatedBy = item.CreatedBy
	updated.CreatedAt = item.CreatedAt
	updated.UpdatedBy = auth.UserID
	updated.DeletedAt = item.DeletedAt
	updated, err = u.repo.UpdateValue(updated)
	if err != nil {
		return model.SKUOptionValueResponse{}, err
	}
	return mapOptionValue(updated, lang), nil
}

func (u *SKUOptionUsecase) DeleteValue(auth model.AuthContext, id string) error {
	return u.repo.DeleteValue(auth.MerchantID, id)
}

func validateOptionGroupRequest(req *model.SKUOptionGroupRequest) error {
	selection := strings.ToUpper(strings.TrimSpace(req.SelectionType))
	switch selection {
	case OptionSelectionRadio:
		req.SelectionType = OptionSelectionRadio
		req.MinSelect = clampMin(req.MinSelect, 0)
		if req.IsRequired && req.MinSelect == 0 {
			req.MinSelect = 1
		}
		req.MaxSelect = 1
	case OptionSelectionCheckbox:
		req.SelectionType = OptionSelectionCheckbox
		req.MinSelect = clampMin(req.MinSelect, 0)
		if req.MaxSelect < 0 {
			req.MaxSelect = 0
		}
		if req.MaxSelect > 0 && req.MinSelect > req.MaxSelect {
			return errors.New("min_select cannot exceed max_select")
		}
	default:
		return errors.New("selection_type must be RADIO or CHECKBOX")
	}
	return nil
}

func clampMin(value, min int) int {
	if value < min {
		return min
	}
	return value
}

func optionGroupFromRequest(auth model.AuthContext, skuID string, req model.SKUOptionGroupRequest) entity.SKUOptionGroup {
	item := entity.SKUOptionGroup{Base: entity.Base{MerchantID: auth.MerchantID, CreatedBy: auth.UserID, UpdatedBy: auth.UserID}, SKUID: skuID}
	applyOptionGroupRequest(&item, req)
	return item
}

func applyOptionGroupRequest(item *entity.SKUOptionGroup, req model.SKUOptionGroupRequest) {
	item.Code = req.Code
	item.NameTH = req.NameTH
	item.NameEN = req.NameEN
	item.SelectionType = strings.ToUpper(req.SelectionType)
	item.IsRequired = req.IsRequired
	item.MinSelect = req.MinSelect
	item.MaxSelect = req.MaxSelect
	item.SortOrder = req.SortOrder
	item.IsActive = boolPointerDefault(req.IsActive, true)
}

func optionValuesFromRequest(items []model.SKUOptionValueRequest) []entity.SKUOptionValue {
	out := make([]entity.SKUOptionValue, 0, len(items))
	for _, item := range items {
		out = append(out, optionValueFromRequest(item))
	}
	return out
}

func optionValueFromRequest(req model.SKUOptionValueRequest) entity.SKUOptionValue {
	return entity.SKUOptionValue{Code: req.Code, NameTH: req.NameTH, NameEN: req.NameEN, PriceDelta: req.PriceDelta, SortOrder: req.SortOrder, IsDefault: req.IsDefault, IsActive: boolPointerDefault(req.IsActive, true)}
}

func mapOptionGroup(item entity.SKUOptionGroup, lang string) model.SKUOptionGroupResponse {
	options := make([]model.SKUOptionValueResponse, 0, len(item.Options))
	for _, option := range item.Options {
		options = append(options, mapOptionValue(option, lang))
	}
	name := item.NameEN
	if languageCode(lang) == "th" {
		name = item.NameTH
	}
	return model.SKUOptionGroupResponse{ID: item.ID.String(), MerchantID: item.MerchantID, SKUID: item.SKUID, Code: item.Code, NameTH: item.NameTH, NameEN: item.NameEN, Name: name, SelectionType: item.SelectionType, IsRequired: item.IsRequired, MinSelect: item.MinSelect, MaxSelect: item.MaxSelect, SortOrder: item.SortOrder, IsActive: item.IsActive, Options: options, AuditFields: model.AuditFields{CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt, CreatedBy: item.CreatedBy, UpdatedBy: item.UpdatedBy}}
}

func mapOptionValue(item entity.SKUOptionValue, lang string) model.SKUOptionValueResponse {
	name := item.NameEN
	if languageCode(lang) == "th" {
		name = item.NameTH
	}
	return model.SKUOptionValueResponse{ID: item.ID.String(), MerchantID: item.MerchantID, OptionGroupID: item.OptionGroupID, Code: item.Code, NameTH: item.NameTH, NameEN: item.NameEN, Name: name, PriceDelta: item.PriceDelta, SortOrder: item.SortOrder, IsDefault: item.IsDefault, IsActive: item.IsActive, AuditFields: model.AuditFields{CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt, CreatedBy: item.CreatedBy, UpdatedBy: item.UpdatedBy}}
}

func boolPointerDefault(value *bool, fallback bool) bool {
	if value == nil {
		return fallback
	}
	return *value
}
