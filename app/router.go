package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupRouter(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate, logger *logrus.Logger) {
	// Inisialisasi Repositories
	// exampleRepo := repositories.NewExampleRepository(db)

	// Inisialisasi Services
	// exampleService := services.NewExampleService(exampleRepo, validate)

	// Inisialisasi Controllers
	// exampleController := controllers.NewExampleController(exampleService, logger)

	api := fiberApp.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Stokeria service is running",
		})
	})

	// 4. Daftarkan Routes
	// api.Post("/example", exampleController.Create)
}