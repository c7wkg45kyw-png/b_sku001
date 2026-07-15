package handler

import (
	"net/http"

	"bsku001/backend/internal/middleware"
	"bsku001/backend/internal/model"
	"bsku001/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ConfigHandler struct {
	usecase *usecase.ConfigUsecase
}

func NewConfigHandler(usecase *usecase.ConfigUsecase) *ConfigHandler {
	return &ConfigHandler{usecase: usecase}
}

func (h *ConfigHandler) GetMerchantFeatureConfig(c *gin.Context) {
	result, err := h.usecase.GetMerchantFeatureConfig(middleware.CurrentAuth(c))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "merchant feature config", result)
}

func (h *ConfigHandler) UpdateMerchantFeatureConfig(c *gin.Context) {
	var req model.MerchantFeatureConfigUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.UpdateMerchantFeatureConfig(middleware.CurrentAuth(c), req)
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "merchant feature config updated", result)
}
