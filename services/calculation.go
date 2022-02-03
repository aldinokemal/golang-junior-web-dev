package services

import (
	"github.com/gofiber/fiber/v2"
	"lsp/structs"
)

type CalculationService interface {
	CalculatePersegi(c *fiber.Ctx, request structs.CalculatePersegiRequest) (response structs.CalculatePersegiResponse, err error)
	HistoryPersegi(c *fiber.Ctx) (response structs.HistoryPersegi, err error)

	CalculateLingkaran(c *fiber.Ctx, request structs.CalculateLingkaranRequest) (response structs.CalculateLingkaranResponse, err error)
	HistoryLingkaran(c *fiber.Ctx) (response structs.HistoryLingkaran, err error)

	CalculateSegitiga(c *fiber.Ctx, request structs.CalculateSegitigaRequest) (response structs.CalculateSegitigaResponse, err error)
	HistorySegitiga(c *fiber.Ctx) (response structs.HistorySegitiga, err error)
}
