package usecase

import (
	"encoding/json"
	"math"
	"strings"

	"bsku001/backend/internal/entity"
	"bsku001/backend/internal/model"
	"bsku001/backend/internal/repository"

	"gorm.io/datatypes"
)

type SKUUsecase struct{ repo *repository.SKURepository }

func NewSKUUsecase(repo *repository.SKURepository) *SKUUsecase { return &SKUUsecase{repo: repo} }

func (u *SKUUsecase) List(auth model.AuthContext, query model.SKUListQuery, lang string) (model.PageResponse[model.SKUResponse], error) {
	items, total, page, limit, err := u.repo.List(auth.MerchantID, query)
	if err != nil {
		return model.PageResponse[model.SKUResponse]{}, err
	}
	out := make([]model.SKUResponse, 0, len(items))
	for _, item := range items {
		out = append(out, mapSKU(item, lang))
	}
	return model.PageResponse[model.SKUResponse]{Items: out, Page: page, Limit: limit, Total: total, TotalPages: int(math.Ceil(float64(total) / float64(limit)))}, nil
}

func (u *SKUUsecase) Get(auth model.AuthContext, idOrCode, lang string) (model.SKUResponse, error) {
	item, err := u.repo.Get(auth.MerchantID, idOrCode)
	if err != nil {
		return model.SKUResponse{}, err
	}
	return mapSKU(item, lang), nil
}

func (u *SKUUsecase) Create(auth model.AuthContext, req model.SKURequest, lang string) (model.SKUResponse, error) {
	item := entity.SKU{Base: entity.Base{MerchantID: auth.MerchantID, CreatedBy: auth.UserID, UpdatedBy: auth.UserID}, SKUCode: req.SKUCode, GTIN: req.GTIN, Names: jsonMap(req.Names), Descriptions: jsonMap(nil), BrandID: optionalString(req.BrandID), CategoryID: optionalString(req.CategoryID), SubCategoryID: optionalString(req.SubCategoryID), MaterialID: optionalString(req.MaterialID), ColorID: optionalString(req.ColorID), Status: statusDefault(req.Status), IsHazmat: req.IsHazmat, CountryOfOrigin: strings.ToUpper(req.CountryOfOrigin), IsActive: true}
	dimension := dimensionFromRequest(req.Weight, req.Dimensions)
	images := imagesFromRequest(req.Images)
	created, err := u.repo.Create(item, dimension, images)
	if err != nil {
		return model.SKUResponse{}, err
	}
	return mapSKU(created, lang), nil
}

func (u *SKUUsecase) Replace(auth model.AuthContext, id string, req model.SKURequest, lang string) (model.SKUResponse, error) {
	item, err := u.repo.Get(auth.MerchantID, id)
	if err != nil {
		return model.SKUResponse{}, err
	}
	item.SKUCode = req.SKUCode
	item.GTIN = req.GTIN
	item.Names = jsonMap(req.Names)
	item.Descriptions = jsonMap(nil)
	item.BrandID = optionalString(req.BrandID)
	item.CategoryID = optionalString(req.CategoryID)
	item.SubCategoryID = optionalString(req.SubCategoryID)
	item.MaterialID = optionalString(req.MaterialID)
	item.ColorID = optionalString(req.ColorID)
	item.Status = statusDefault(req.Status)
	item.IsHazmat = req.IsHazmat
	item.CountryOfOrigin = strings.ToUpper(req.CountryOfOrigin)
	item.UpdatedBy = auth.UserID
	updated, err := u.repo.Update(item, dimensionFromRequest(req.Weight, req.Dimensions), true, imagesFromRequest(req.Images))
	if err != nil {
		return model.SKUResponse{}, err
	}
	return mapSKU(updated, lang), nil
}

func (u *SKUUsecase) Patch(auth model.AuthContext, id string, req model.SKUPatchRequest, lang string) (model.SKUResponse, error) {
	item, err := u.repo.Get(auth.MerchantID, id)
	if err != nil {
		return model.SKUResponse{}, err
	}
	if req.SKUCode != nil {
		item.SKUCode = *req.SKUCode
	}
	if req.GTIN != nil {
		item.GTIN = *req.GTIN
	}
	if req.Names != nil {
		item.Names = jsonMap(req.Names)
	}
	if req.BrandID != nil {
		item.BrandID = optionalString(req.BrandID)
	}
	if req.CategoryID != nil {
		item.CategoryID = optionalString(req.CategoryID)
	}
	if req.SubCategoryID != nil {
		item.SubCategoryID = optionalString(req.SubCategoryID)
	}
	if req.MaterialID != nil {
		item.MaterialID = optionalString(req.MaterialID)
	}
	if req.ColorID != nil {
		item.ColorID = optionalString(req.ColorID)
	}
	if req.Status != nil {
		item.Status = statusDefault(*req.Status)
	}
	if req.IsHazmat != nil {
		item.IsHazmat = *req.IsHazmat
	}
	if req.CountryOfOrigin != nil {
		item.CountryOfOrigin = strings.ToUpper(*req.CountryOfOrigin)
	}
	if req.IsActive != nil {
		item.IsActive = *req.IsActive
	}
	item.UpdatedBy = auth.UserID
	dimension := item.Dimension
	if req.Weight != nil {
		dimension.Weight = req.Weight.Value
		dimension.WeightUnit = unitDefault(req.Weight.Unit, "kg")
	}
	if req.Dimensions != nil {
		applyDimension(&dimension, *req.Dimensions)
	}
	updated, err := u.repo.Update(item, dimension, false, nil)
	if err != nil {
		return model.SKUResponse{}, err
	}
	return mapSKU(updated, lang), nil
}

func (u *SKUUsecase) Delete(auth model.AuthContext, id string) error {
	return u.repo.SoftDelete(auth.MerchantID, id)
}

func optionalString(value *string) *string {
	if value == nil || strings.TrimSpace(*value) == "" {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	return &trimmed
}

func jsonMap(value map[string]string) datatypes.JSON {
	if value == nil {
		value = map[string]string{}
	}
	raw, _ := json.Marshal(value)
	return datatypes.JSON(raw)
}
func statusDefault(value string) string {
	if strings.TrimSpace(value) == "" {
		return "DRAFT"
	}
	return strings.ToUpper(value)
}
func unitDefault(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}
func dimensionFromRequest(weight model.WeightPayload, dimensions model.DimensionPayload) entity.SKUDimension {
	d := entity.SKUDimension{Weight: weight.Value, WeightUnit: unitDefault(weight.Unit, "kg")}
	applyDimension(&d, dimensions)
	return d
}
func applyDimension(d *entity.SKUDimension, dimensions model.DimensionPayload) {
	d.Width = dimensions.Width.Value
	d.Length = dimensions.Length.Value
	d.Height = dimensions.Height.Value
	d.Unit = unitDefault(dimensions.Width.Unit, unitDefault(dimensions.Length.Unit, unitDefault(dimensions.Height.Unit, "cm")))
	d.HSCode = dimensions.HSCode
}
func imagesFromRequest(items []model.SKUImageRequest) []entity.SKUImage {
	out := make([]entity.SKUImage, 0, len(items))
	for _, item := range items {
		out = append(out, entity.SKUImage{ImageURL: item.ImageURL, AltText: item.AltText, SortOrder: item.SortOrder, IsPrimary: item.IsPrimary})
	}
	return out
}

func mapSKU(item entity.SKU, lang string) model.SKUResponse {
	names := mapString(item.Names)
	descriptions := mapString(item.Descriptions)
	code := languageCode(lang)
	images := make([]model.SKUImageResponse, 0, len(item.Images))
	for _, image := range item.Images {
		images = append(images, model.SKUImageResponse{ID: image.ID.String(), ImageURL: image.ImageURL, AltText: image.AltText, SortOrder: image.SortOrder, IsPrimary: image.IsPrimary})
	}
	return model.SKUResponse{ID: item.ID.String(), MerchantID: item.MerchantID, SKUCode: item.SKUCode, GTIN: item.GTIN, BrandID: item.BrandID, CategoryID: item.CategoryID, SubCategoryID: item.SubCategoryID, MaterialID: item.MaterialID, ColorID: item.ColorID, Names: names, Name: names[code], Descriptions: descriptions, Description: descriptions[code], Weight: model.WeightPayload{Value: item.Dimension.Weight, Unit: item.Dimension.WeightUnit}, Dimensions: model.DimensionPayload{Width: model.UnitValue{Value: item.Dimension.Width, Unit: item.Dimension.Unit}, Length: model.UnitValue{Value: item.Dimension.Length, Unit: item.Dimension.Unit}, Height: model.UnitValue{Value: item.Dimension.Height, Unit: item.Dimension.Unit}, HSCode: item.Dimension.HSCode}, HSCode: item.Dimension.HSCode, Status: item.Status, IsHazmat: item.IsHazmat, CountryOfOrigin: item.CountryOfOrigin, IsActive: item.IsActive, Images: images, AuditFields: model.AuditFields{CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt, CreatedBy: item.CreatedBy, UpdatedBy: item.UpdatedBy}}
}
func mapString(raw datatypes.JSON) map[string]string {
	out := map[string]string{}
	_ = json.Unmarshal(raw, &out)
	return out
}
func languageCode(header string) string {
	lower := strings.ToLower(header)
	if strings.HasPrefix(lower, "th") {
		return "th"
	}
	return "en"
}

func (u *SKUUsecase) GetDimension(auth model.AuthContext, skuID string) (model.SKUDimensionResponse, error) {
	item, err := u.repo.GetDimension(auth.MerchantID, skuID)
	if err != nil {
		return model.SKUDimensionResponse{}, err
	}
	return mapDimension(item), nil
}

func (u *SKUUsecase) UpsertDimension(auth model.AuthContext, skuID string, req model.SKUDimensionRequest) (model.SKUDimensionResponse, error) {
	item := entity.SKUDimension{Base: entity.Base{MerchantID: auth.MerchantID, CreatedBy: auth.UserID, UpdatedBy: auth.UserID}, SKUID: skuID, Weight: req.Weight, WeightUnit: unitDefault(req.WeightUnit, "kg"), Width: req.Width, Length: req.Length, Height: req.Height, Unit: unitDefault(req.Unit, "cm"), HSCode: req.HSCode}
	created, err := u.repo.UpsertDimension(item)
	if err != nil {
		return model.SKUDimensionResponse{}, err
	}
	return mapDimension(created), nil
}

func (u *SKUUsecase) ListImages(auth model.AuthContext, skuID string) ([]model.SKUImageResponse, error) {
	items, err := u.repo.ListImages(auth.MerchantID, skuID)
	if err != nil {
		return nil, err
	}
	out := make([]model.SKUImageResponse, 0, len(items))
	for _, item := range items {
		out = append(out, mapImage(item))
	}
	return out, nil
}

func (u *SKUUsecase) CreateImage(auth model.AuthContext, skuID string, req model.SKUImageRequest) (model.SKUImageResponse, error) {
	item := entity.SKUImage{Base: entity.Base{MerchantID: auth.MerchantID, CreatedBy: auth.UserID, UpdatedBy: auth.UserID}, SKUID: skuID, ImageURL: req.ImageURL, AltText: req.AltText, SortOrder: req.SortOrder, IsPrimary: req.IsPrimary}
	created, err := u.repo.CreateImage(item)
	if err != nil {
		return model.SKUImageResponse{}, err
	}
	return mapImage(created), nil
}

func (u *SKUUsecase) UpdateImage(auth model.AuthContext, imageID string, req model.SKUImageRequest) (model.SKUImageResponse, error) {
	item := entity.SKUImage{Base: entity.Base{UpdatedBy: auth.UserID}, ImageURL: req.ImageURL, AltText: req.AltText, SortOrder: req.SortOrder, IsPrimary: req.IsPrimary}
	updated, err := u.repo.UpdateImage(auth.MerchantID, imageID, item)
	if err != nil {
		return model.SKUImageResponse{}, err
	}
	return mapImage(updated), nil
}

func (u *SKUUsecase) DeleteImage(auth model.AuthContext, imageID string) error {
	return u.repo.DeleteImage(auth.MerchantID, imageID)
}

func mapDimension(item entity.SKUDimension) model.SKUDimensionResponse {
	return model.SKUDimensionResponse{ID: item.ID.String(), SKUID: item.SKUID, Weight: item.Weight, WeightUnit: item.WeightUnit, Width: item.Width, Length: item.Length, Height: item.Height, Unit: item.Unit, HSCode: item.HSCode, AuditFields: model.AuditFields{CreatedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt, CreatedBy: item.CreatedBy, UpdatedBy: item.UpdatedBy}}
}

func mapImage(item entity.SKUImage) model.SKUImageResponse {
	return model.SKUImageResponse{ID: item.ID.String(), ImageURL: item.ImageURL, AltText: item.AltText, SortOrder: item.SortOrder, IsPrimary: item.IsPrimary}
}
