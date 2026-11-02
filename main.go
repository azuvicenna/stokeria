package main

import (
	"stokeria/app"

	"github.com/gofiber/fiber/v2"
)

func main(){
	viperConfig := app.NewViper()
	logrusLogger := app.NewLogrus()
	gormDB := app.NewDatabase(viperConfig, logrusLogger)
	goValidator := app.NewValidator()

	fiberApp := fiber.New(fiber.Config{
		// Tambahkan custom error handler jika didefinisikan di direktori exception
		// ErrorHandler: exception.ErrorHandler,
	})

	app.SetupRouter(fiberApp, gormDB, goValidator, logrusLogger)

	appPort := viperConfig.GetString("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	err := fiberApp.Listen(":" + appPort)
	if err != nil {
		logrusLogger.Fatalf("Server Failed to Start: %v", err)
	}
}