package handler

import (
	"net/http"

	"bsku001/backend/internal/middleware"
	"bsku001/backend/internal/model"
	"bsku001/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SKUOptionHandler struct{ usecase *usecase.SKUOptionUsecase }

func NewSKUOptionHandler(usecase *usecase.SKUOptionUsecase) *SKUOptionHandler {
	return &SKUOptionHandler{usecase: usecase}
}

func (h *SKUOptionHandler) ListGroups(c *gin.Context) {
	var query model.SKUOptionGroupListQuery
	_ = c.ShouldBindQuery(&query)
	result, err := h.usecase.ListGroups(middleware.CurrentAuth(c), query, c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku option groups", result)
}

func (h *SKUOptionHandler) ListGroupsBySKU(c *gin.Context) {
	result, err := h.usecase.ListGroupsBySKU(middleware.CurrentAuth(c), skuParam(c), c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku options", result)
}

func (h *SKUOptionHandler) GetGroup(c *gin.Context) {
	result, err := h.usecase.GetGroup(middleware.CurrentAuth(c), c.Param("id"), c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku option group", result)
}

func (h *SKUOptionHandler) CreateGroup(c *gin.Context) {
	var req model.SKUOptionGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.CreateGroup(middleware.CurrentAuth(c), skuParam(c), req, c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	created(c, "sku option group created", result)
}

func (h *SKUOptionHandler) ReplaceGroup(c *gin.Context) {
	var req model.SKUOptionGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.ReplaceGroup(middleware.CurrentAuth(c), c.Param("id"), req, c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku option group updated", result)
}

func (h *SKUOptionHandler) DeleteGroup(c *gin.Context) {
	if err := h.usecase.DeleteGroup(middleware.CurrentAuth(c), c.Param("id")); err != nil {
		handleError(c, err)
		return
	}
	noContent(c)
}

func (h *SKUOptionHandler) CreateValue(c *gin.Context) {
	var req model.SKUOptionValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.CreateValue(middleware.CurrentAuth(c), c.Param("id"), req, c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	created(c, "sku option value created", result)
}

func (h *SKUOptionHandler) ReplaceValue(c *gin.Context) {
	var req model.SKUOptionValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.ReplaceValue(middleware.CurrentAuth(c), c.Param("id"), req, c.GetHeader("Accept-Language"))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "sku option value updated", result)
}

func (h *SKUOptionHandler) DeleteValue(c *gin.Context) {
	if err := h.usecase.DeleteValue(middleware.CurrentAuth(c), c.Param("id")); err != nil {
		handleError(c, err)
		return
	}
	noContent(c)
}
