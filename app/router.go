package app

import (
	"stokeria/controllers"
	"stokeria/repository"
	"stokeria/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupRouter(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate, logger *logrus.Logger) {
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
	userService := service.NewUserService(userRepo, validate)
	supplierService := service.NewSupplierService(supplierRepo, validate)
	warehouseService := service.NewWarehouseService(warehouseRepo, validate)
	productService := service.NewProductService(productRepo, validate)
	inventoryService := service.NewInventoryService(inventoryRepo)
	inboundService := service.NewInboundTransactionService(inboundRepo, detailRepo, validate)
	outboundService := service.NewOutboundTransactionService(outboundRepo, detailRepo, validate)
	stockAdjustmentService := service.NewStockAdjustmentService(stockAdjustmentRepo, validate)

	// 3. Inisialisasi Controllers & Registrasi Routes
	categoryController := controller.NewCategoryController(categoryService)
	categoryController.Route(fiberApp)

	userController := controller.NewUserController(userService)
	userController.Route(fiberApp)

	supplierController := controller.NewSupplierController(supplierService)
	supplierController.Route(fiberApp)

	warehouseController := controller.NewWarehouseController(warehouseService)
	warehouseController.Route(fiberApp)

	productController := controller.NewProductController(productService)
	productController.Route(fiberApp)

	inventoryController := controller.NewInventoryController(inventoryService)
	inventoryController.Route(fiberApp)

	inboundController := controller.NewInboundTransactionController(inboundService)
	inboundController.Route(fiberApp)

	outboundController := controller.NewOutboundTransactionController(outboundService)
	outboundController.Route(fiberApp)

	stockAdjustmentController := controller.NewStockAdjustmentController(stockAdjustmentService)
	stockAdjustmentController.Route(fiberApp)

	// 4. Health Check Route
	api := fiberApp.Group("/api")
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Stokeria service is running",
		})
	})

	logger.Info("Semua repository, service, dan controller berhasil didaftarkan ke router.")
}
