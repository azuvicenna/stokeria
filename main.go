package main

import (
	"flag"
	"stokeria/app"
	"stokeria/exception"

	"github.com/gofiber/fiber/v2"
)

func main(){
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