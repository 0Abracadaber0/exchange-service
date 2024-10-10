package router

import (
	"exchange/internal/config"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetRatesHandler(ctx *fiber.Ctx) error {
	return nil
}

func GetRatesInDayHandler(ctx *fiber.Ctx) error {
	return nil
}

func SetupRoutes(app *fiber.App, cfg *config.Config, log *slog.Logger) {

	app.Use(func(ctx *fiber.Ctx) error {
		start := time.Now()

		err := ctx.Next()

		log.Info("HTTP request",
			slog.String("method", ctx.Method()),
			slog.String("route", ctx.Path()),
			slog.Int("status", ctx.Response().StatusCode()),
			slog.Duration("latency", time.Since(start)),
		)

		return err
	})

	app.Get("/rates", GetRatesHandler)
	app.Get("/rates/:day", GetRatesInDayHandler)
}
