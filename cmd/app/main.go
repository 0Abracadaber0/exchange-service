package main

import (
	"exchange/internal/config"
	"exchange/internal/database"
	"exchange/internal/router"
	"exchange/internal/service"
	"log/slog"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.MustLoad()

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log.Info("app is starting...", "cfg", cfg)

	app := fiber.New(fiber.Config{
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	})

	if err := database.ConnectDB(log, cfg); err != nil {
		panic(err)
	}
	log.Info("succesfull connection to the database")
	if err := database.RunMigrations(log, cfg); err != nil {
		panic(err)
	}
	log.Info("succesfull migrations")

	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	scheduler := gocron.NewScheduler(location)
	scheduler.Every(1).Day().At("15:38").Do(service.FetchAndStoreData, log, cfg)

	scheduler.StartAsync()
	log.Debug("scheduler started")

	router.SetupRoutes(app, log)

	if err := app.Listen(cfg.AppHost.Value + ":" + cfg.AppPort.Value); err != nil {
		panic(err)
	}
}
