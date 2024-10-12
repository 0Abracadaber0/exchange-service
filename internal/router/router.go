package router

import (
	"exchange/internal/handlers"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "exchange/internal/models"

	_ "exchange/docs"
)

// GetRatesHandler godoc
// @Summary Get all currency rates
// @Description Get a list of all currency exchange rates.
// @Tags rates
// @Produce json
// @Success 200 {array} models.ExchangeRate
// @Failure 500 {object} map[string]interface{}
// @Router /rates [get]
func GetRatesHandler(ctx *fiber.Ctx) error {
	return handlers.GetRatesHandler(ctx)
}

// GetRatesByDateHandler godoc
// @Summary Get currency rates by date
// @Description Get currency exchange rates for a specific date in YYYY-MM-DD format.
// @Tags rates
// @Produce json
// @Param date path string true "Date in YYYY-MM-DD format"
// @Success 200 {array} models.ExchangeRate
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rates/{date} [get]
func GetRatesByDateHandler(ctx *fiber.Ctx) error {
	return handlers.GetRatesByDateHandler(ctx)
}

func SetupRoutes(app *fiber.App, log *slog.Logger) {
	log.Debug("setting up routes")

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

	// Swagger UI route
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/rates", GetRatesHandler)
	app.Get("/rates/:date", GetRatesByDateHandler)
}
