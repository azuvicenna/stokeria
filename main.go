package main

import (
	"flag"
	"stokeria/app"
	"stokeria/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	viperConfig := app.NewViper()
	logrusLogger := app.NewLogrus()
	gormDB := app.NewDatabase(viperConfig, logrusLogger)
	goValidator := app.NewValidator()

	seedFlag := flag.Bool("seed", false, "Menjalankan database seeder")
	flag.Parse()

	if *seedFlag {
		if err := app.RunSeeder(gormDB); err != nil {
			logrusLogger.Fatalf("Seeder gagal dijalankan: %v", err)
		}
		logrusLogger.Info("Seeder berhasil dieksekusi!")
		return
	}

	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	// Middleware Global
	fiberApp.Use(recover.New())
	fiberApp.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} | ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	app.SetupRouter(fiberApp, gormDB, goValidator, logrusLogger, viperConfig)

	appPort := viperConfig.GetString("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	err := fiberApp.Listen(":" + appPort)
	if err != nil {
		logrusLogger.Fatalf("Server Failed to Start: %v", err)
	}
}
