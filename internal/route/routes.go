package route

import (
	"net/http"

	"bsku001/backend/internal/config"
	"bsku001/backend/internal/handler"
	"bsku001/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	SKU     *handler.SKUHandler
	Master  *handler.MasterHandler
	Config  *handler.ConfigHandler
	Options *handler.SKUOptionHandler
	Summary *handler.SummaryHandler
}

func Register(router *gin.Engine, cfg config.Config, handlers Handlers) {
	router.Use(middleware.CORS(cfg))
	router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "message": "healthy"}) })
	router.StaticFile("/docs/openapi.yaml", "./docs/openapi.yaml")
	router.GET("/docs/swagger.html", swaggerHTML)

	api := router.Group("/api/v1")
	api.Use(middleware.Auth(cfg))

	api.GET("/merchant-config", middleware.RequireScope("sku:read"), handlers.Config.GetMerchantFeatureConfig)
	api.PATCH("/merchant-config", middleware.RequireScope("sku:update"), handlers.Config.UpdateMerchantFeatureConfig)
	api.GET("/summary", middleware.RequireScope("sku:read"), handlers.Summary.Get)

	api.GET("/skus", middleware.RequireScope("sku:read"), handlers.SKU.List)
	api.POST("/skus", middleware.RequireScope("sku:create"), handlers.SKU.Create)
	api.GET("/skus/:id_or_code", middleware.RequireScope("sku:read"), handlers.SKU.Get)
	api.PUT("/skus/:id_or_code", middleware.RequireScope("sku:update"), handlers.SKU.Replace)
	api.PATCH("/skus/:id_or_code", middleware.RequireScope("sku:update"), handlers.SKU.Patch)
	api.DELETE("/skus/:id_or_code", middleware.RequireScope("sku:delete"), handlers.SKU.Delete)
	api.GET("/skus/:id_or_code/dimensions", middleware.RequireScope("sku:read"), handlers.SKU.GetDimension)
	api.PUT("/skus/:id_or_code/dimensions", middleware.RequireScope("sku:update"), handlers.SKU.UpsertDimension)
	api.PATCH("/skus/:id_or_code/dimensions", middleware.RequireScope("sku:update"), handlers.SKU.UpsertDimension)
	api.GET("/skus/:id_or_code/images", middleware.RequireScope("sku:read"), handlers.SKU.ListImages)
	api.POST("/skus/:id_or_code/images", middleware.RequireScope("sku:create"), handlers.SKU.CreateImage)
	api.PUT("/sku-images/:image_id", middleware.RequireScope("sku:update"), handlers.SKU.UpdateImage)
	api.PATCH("/sku-images/:image_id", middleware.RequireScope("sku:update"), handlers.SKU.UpdateImage)
	api.DELETE("/sku-images/:image_id", middleware.RequireScope("sku:delete"), handlers.SKU.DeleteImage)

	api.GET("/sku-option-groups", middleware.RequireScope("sku:read"), handlers.Options.ListGroups)
	api.GET("/sku-option-groups/:id", middleware.RequireScope("sku:read"), handlers.Options.GetGroup)
	api.PUT("/sku-option-groups/:id", middleware.RequireScope("sku:update"), handlers.Options.ReplaceGroup)
	api.PATCH("/sku-option-groups/:id", middleware.RequireScope("sku:update"), handlers.Options.ReplaceGroup)
	api.DELETE("/sku-option-groups/:id", middleware.RequireScope("sku:delete"), handlers.Options.DeleteGroup)
	api.POST("/sku-option-groups/:id/options", middleware.RequireScope("sku:create"), handlers.Options.CreateValue)
	api.PUT("/sku-option-values/:id", middleware.RequireScope("sku:update"), handlers.Options.ReplaceValue)
	api.PATCH("/sku-option-values/:id", middleware.RequireScope("sku:update"), handlers.Options.ReplaceValue)
	api.DELETE("/sku-option-values/:id", middleware.RequireScope("sku:delete"), handlers.Options.DeleteValue)
	api.GET("/skus/:id_or_code/options", middleware.RequireScope("sku:read"), handlers.Options.ListGroupsBySKU)
	api.POST("/skus/:id_or_code/options", middleware.RequireScope("sku:create"), handlers.Options.CreateGroup)

	registerMaster(api, "brands", handlers.Master)
	registerMaster(api, "categories", handlers.Master)
	registerMaster(api, "sub-categories", handlers.Master)
	registerMaster(api, "materials", handlers.Master)
	registerMaster(api, "colors", handlers.Master)
}

func registerMaster(api *gin.RouterGroup, resource string, h *handler.MasterHandler) {
	base := "/" + resource
	api.GET(base, middleware.RequireScope("sku:read"), h.List(resource))
	api.POST(base, middleware.RequireScope("sku:create"), h.Create(resource))
	api.GET(base+"/:id", middleware.RequireScope("sku:read"), h.Get(resource))
	api.PUT(base+"/:id", middleware.RequireScope("sku:update"), h.Update(resource))
	api.PATCH(base+"/:id", middleware.RequireScope("sku:update"), h.Update(resource))
	api.DELETE(base+"/:id", middleware.RequireScope("sku:delete"), h.Delete(resource))
}

func swaggerHTML(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, `<!doctype html><html><head><title>BSKU001 Swagger</title><link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css"></head><body><div id="swagger-ui"></div><script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script><script>window.onload=()=>SwaggerUIBundle({url:'/docs/openapi.yaml',dom_id:'#swagger-ui'});</script></body></html>`)
}
