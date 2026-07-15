package handler

import (
	"net/http"

	"bsku001/backend/internal/middleware"
	"bsku001/backend/internal/model"
	"bsku001/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MasterHandler struct{ usecase *usecase.MasterUsecase }

func NewMasterHandler(usecase *usecase.MasterUsecase) *MasterHandler {
	return &MasterHandler{usecase: usecase}
}

func (h *MasterHandler) List(resource string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query model.ListQuery
		_ = c.ShouldBindQuery(&query)
		result, err := h.usecase.List(resource, middleware.CurrentAuth(c), query)
		if err != nil {
			handleError(c, err)
			return
		}
		ok(c, resource, result)
	}
}
func (h *MasterHandler) Get(resource string) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := h.usecase.Get(resource, middleware.CurrentAuth(c), c.Param("id"))
		if err != nil {
			handleError(c, err)
			return
		}
		ok(c, resource, result)
	}
}
func (h *MasterHandler) Create(resource string) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, bindOK := bindMasterRequest(c, resource)
		if !bindOK {
			return
		}
		result, err := h.usecase.Create(resource, middleware.CurrentAuth(c), req)
		if err != nil {
			handleError(c, err)
			return
		}
		created(c, resource+" created", result)
	}
}
func (h *MasterHandler) Update(resource string) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, bindOK := bindMasterRequest(c, resource)
		if !bindOK {
			return
		}
		result, err := h.usecase.Update(resource, middleware.CurrentAuth(c), c.Param("id"), req)
		if err != nil {
			handleError(c, err)
			return
		}
		ok(c, resource+" updated", result)
	}
}
func (h *MasterHandler) Delete(resource string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h.usecase.Delete(resource, middleware.CurrentAuth(c), c.Param("id")); err != nil {
			handleError(c, err)
			return
		}
		noContent(c)
	}
}

func bindMasterRequest(c *gin.Context, resource string) (any, bool) {
	if resource == "sub-categories" {
		var req model.SubCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
			return nil, false
		}
		return req, true
	}
	var req model.MasterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return nil, false
	}
	return req, true
}
