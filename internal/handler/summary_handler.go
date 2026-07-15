package handler

import (
	"bsku001/backend/internal/middleware"
	"bsku001/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SummaryHandler struct{ usecase *usecase.SummaryUsecase }

func NewSummaryHandler(usecase *usecase.SummaryUsecase) *SummaryHandler {
	return &SummaryHandler{usecase: usecase}
}

func (h *SummaryHandler) Get(c *gin.Context) {
	result, err := h.usecase.Get(middleware.CurrentAuth(c))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "summary", result)
}
