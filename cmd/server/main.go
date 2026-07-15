package main

import (
	"bsku001/backend/internal/config"
	"bsku001/backend/internal/database"
	"bsku001/backend/internal/handler"
	"bsku001/backend/internal/logger"
	"bsku001/backend/internal/repository"
	"bsku001/backend/internal/route"
	"bsku001/backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.AppEnv)
	defer log.Sync()

	db, err := database.Connect(cfg.DatabaseDSN)
	if err != nil {
		log.Fatal("connect database", zap.Error(err))
	}
	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("migrate database", zap.Error(err))
	}

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	masterRepo := repository.NewMasterRepository(db)
	skuRepo := repository.NewSKURepository(db)
	skuOptionRepo := repository.NewSKUOptionRepository(db)
	configRepo := repository.NewConfigRepository(db)
	summaryRepo := repository.NewSummaryRepository(db)
	masterUsecase := usecase.NewMasterUsecase(masterRepo)
	skuUsecase := usecase.NewSKUUsecase(skuRepo)
	skuOptionUsecase := usecase.NewSKUOptionUsecase(skuOptionRepo)
	configUsecase := usecase.NewConfigUsecase(configRepo)
	summaryUsecase := usecase.NewSummaryUsecase(summaryRepo)
	route.Register(r, cfg, route.Handlers{SKU: handler.NewSKUHandler(skuUsecase), Master: handler.NewMasterHandler(masterUsecase), Config: handler.NewConfigHandler(configUsecase), Options: handler.NewSKUOptionHandler(skuOptionUsecase), Summary: handler.NewSummaryHandler(summaryUsecase)})

	log.Info("BSKU001 backend started", zap.String("port", cfg.AppPort))
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("server stopped", zap.Error(err))
	}
}
