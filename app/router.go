package app

import (
	"stokeria/controller"
	"stokeria/middleware"
	"stokeria/repository"
	"stokeria/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func SetupRouter(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate, logger *logrus.Logger, viperConfig *viper.Viper) {
	jwtSecret := "stokeria_secret_jwt_key_2026"
	jwtExp := 24
	if viperConfig != nil {
		if secret := viperConfig.GetString("JWT_SECRET"); secret != "" {
			jwtSecret = secret
		}
		if exp := viperConfig.GetInt("JWT_EXPIRATION_HOURS"); exp > 0 {
			jwtExp = exp
		}
	}

	authMiddleware := middleware.Protected(jwtSecret)

	// 1. Inisialisasi Repositories
	categoryRepo := repository.NewCategoryRepository(db)
	userRepo := repository.NewUserRepository(db)
	supplierRepo := repository.NewSupplierRepository(db)
	warehouseRepo := repository.NewWarehouseRepository(db)
	productRepo := repository.NewProductRepository(db)
	inventoryRepo := repository.NewInventoryRepository(db)
	detailRepo := repository.NewTransactionDetailRepository(db)
	inboundRepo := repository.NewInboundTransactionRepository(db, detailRepo, inventoryRepo)
	outboundRepo := repository.NewOutboundTransactionRepository(db, detailRepo, inventoryRepo)
	stockAdjustmentRepo := repository.NewStockAdjustmentRepository(db, inventoryRepo)

	// 2. Inisialisasi Services
	categoryService := service.NewCategoryService(categoryRepo, validate)
	userService := service.NewUserService(userRepo, validate, jwtSecret, jwtExp)
	supplierService := service.NewSupplierService(supplierRepo, validate)
	warehouseService := service.NewWarehouseService(warehouseRepo, validate)
	productService := service.NewProductService(productRepo, validate)
	inventoryService := service.NewInventoryService(inventoryRepo)
	inboundService := service.NewInboundTransactionService(inboundRepo, detailRepo, validate)
	outboundService := service.NewOutboundTransactionService(outboundRepo, detailRepo, validate)
	stockAdjustmentService := service.NewStockAdjustmentService(stockAdjustmentRepo, validate)

	// 3. Inisialisasi Controllers & Registrasi Routes
	categoryController := controller.NewCategoryController(categoryService)
	categoryController.Route(fiberApp, authMiddleware)

	userController := controller.NewUserController(userService)
	userController.Route(fiberApp, authMiddleware)

	supplierController := controller.NewSupplierController(supplierService)
	supplierController.Route(fiberApp, authMiddleware)

	warehouseController := controller.NewWarehouseController(warehouseService)
	warehouseController.Route(fiberApp, authMiddleware)

	productController := controller.NewProductController(productService)
	productController.Route(fiberApp, authMiddleware)

	inventoryController := controller.NewInventoryController(inventoryService)
	inventoryController.Route(fiberApp, authMiddleware)

	inboundController := controller.NewInboundTransactionController(inboundService)
	inboundController.Route(fiberApp, authMiddleware)

	outboundController := controller.NewOutboundTransactionController(outboundService)
	outboundController.Route(fiberApp, authMiddleware)

	stockAdjustmentController := controller.NewStockAdjustmentController(stockAdjustmentService)
	stockAdjustmentController.Route(fiberApp, authMiddleware)

	// 4. Health Check Route
	api := fiberApp.Group("/api")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Stokeria service is running",
		})
	})

	logger.Info("Semua repository, service, controller, dan auth middleware berhasil didaftarkan ke router.")
}
