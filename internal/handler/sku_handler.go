package handler

import (
	"net/http"

	"bsku001/backend/internal/middleware"
	"bsku001/backend/internal/model"
	"bsku001/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SKUHandler struct{ usecase *usecase.SKUUsecase }

func NewSKUHandler(usecase *usecase.SKUUsecase) *SKUHandler { return &SKUHandler{usecase: usecase} }

func (h *SKUHandler) List(c *gin.Context) {
	var query model.SKUListQuery
	_ = c.ShouldBindQuery(&query)
	result, err := h.usecase.List(middleware.CurrentAuth(c), query, c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "skus", result)
}
func (h *SKUHandler) Get(c *gin.Context) {
	result, err := h.usecase.Get(middleware.CurrentAuth(c), c.Param("id_or_code"), c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku", result)
}
func (h *SKUHandler) Create(c *gin.Context) {
	var req model.SKURequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.Create(middleware.CurrentAuth(c), req, c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	created(c, "sku created", result)
}
func (h *SKUHandler) Replace(c *gin.Context) {
	var req model.SKURequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.Replace(middleware.CurrentAuth(c), skuParam(c), req, c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku updated", result)
}
func (h *SKUHandler) Patch(c *gin.Context) {
	var req model.SKUPatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.Patch(middleware.CurrentAuth(c), skuParam(c), req, c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku patched", result)
}
func (h *SKUHandler) Delete(c *gin.Context) {
	if err := h.usecase.Delete(middleware.CurrentAuth(c), skuParam(c)); err != nil {
		handleError(c, err)
		return
	}
	noContent(c)
}

func (h *SKUHandler) GetDimension(c *gin.Context) {
	result, err := h.usecase.GetDimension(middleware.CurrentAuth(c), skuParam(c))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku dimension", result)
}

func (h *SKUHandler) UpsertDimension(c *gin.Context) {
	var req model.SKUDimensionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.UpsertDimension(middleware.CurrentAuth(c), skuParam(c), req)
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku dimension updated", result)
}

func (h *SKUHandler) ListImages(c *gin.Context) {
	result, err := h.usecase.ListImages(middleware.CurrentAuth(c), skuParam(c))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku images", result)
}

func (h *SKUHandler) CreateImage(c *gin.Context) {
	var req model.SKUImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.CreateImage(middleware.CurrentAuth(c), skuParam(c), req)
	if err != nil {
		handleError(c, err)
		return
	}
	created(c, "sku image created", result)
}

func (h *SKUHandler) UpdateImage(c *gin.Context) {
	var req model.SKUImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.UpdateImage(middleware.CurrentAuth(c), c.Param("image_id"), req)
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku image updated", result)
}

func (h *SKUHandler) DeleteImage(c *gin.Context) {
	if err := h.usecase.DeleteImage(middleware.CurrentAuth(c), c.Param("image_id")); err != nil {
		handleError(c, err)
		return
	}
	noContent(c)
}

func skuParam(c *gin.Context) string {
	if value := c.Param("id_or_code"); value != "" {
		return value
	}
	return c.Param("id")
}
