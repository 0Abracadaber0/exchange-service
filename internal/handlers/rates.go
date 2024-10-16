package handlers

import (
	"exchange/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetRatesHandler(ctx *fiber.Ctx) error {
	rates, err := service.GetAllRates()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(rates)
}

func GetRatesByDateHandler(ctx *fiber.Ctx) error {
	date := ctx.Params("date")
	rates, err := service.GetRatesByDate(date)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(rates)
}
