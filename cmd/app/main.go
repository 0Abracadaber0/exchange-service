package main

import (
	"exchange/internal/config"
	"exchange/internal/database"
	"exchange/internal/router"
	"log/slog"
	"os"
	"time"

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

	router.SetupRoutes(app, log)

	if err := app.Listen(cfg.AppHost.Value + ":" + cfg.AppPort.Value); err != nil {
		panic(err)
	}
}
