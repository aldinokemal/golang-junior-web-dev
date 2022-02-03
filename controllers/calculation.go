package controllers

import (
	"github.com/gofiber/fiber/v2"
	"lsp/services"
	"lsp/structs"
	"lsp/utils"
)

type CalculationController struct {
	Service services.CalculationService
}

func NewCalculationController(service services.CalculationService) CalculationController {
	return CalculationController{Service: service}
}

func (controller *CalculationController) Route(app *fiber.App) {
	app.Get("/calculate/persegi/history", controller.HistoryPersegi)
	app.Post("/calculate/persegi", controller.CalculatePersegi)

	app.Get("/calculate/lingkaran/history", controller.HistoryLingkran)
	app.Post("/calculate/lingkaran", controller.CalculateLingkaran)

	app.Get("/calculate/segitiga/history", controller.HistorySegitiga)
	app.Post("/calculate/segitiga", controller.CalculateSegitiga)
}

func (controller CalculationController) CalculatePersegi(c *fiber.Ctx) error {
	var request structs.CalculatePersegiRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)
	response, err := controller.Service.CalculatePersegi(c, request)
	utils.PanicIfNeeded(err)
	return c.JSON(utils.ResponseData{
		Code:    200,
		Message: response.Message,
		Results: response,
	})
}

func (controller CalculationController) CalculateLingkaran(c *fiber.Ctx) error {
	var request structs.CalculateLingkaranRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)
	response, err := controller.Service.CalculateLingkaran(c, request)
	utils.PanicIfNeeded(err)
	return c.JSON(utils.ResponseData{
		Code:    200,
		Message: response.Message,
		Results: response,
	})
}

func (controller CalculationController) CalculateSegitiga(c *fiber.Ctx) error {
	var request structs.CalculateSegitigaRequest
	err := c.BodyParser(&request)
	utils.PanicIfNeeded(err)
	response, err := controller.Service.CalculateSegitiga(c, request)
	utils.PanicIfNeeded(err)
	return c.JSON(utils.ResponseData{
		Code:    200,
		Message: response.Message,
		Results: response,
	})
}

func (controller CalculationController) HistoryPersegi(c *fiber.Ctx) error {
	response, err := controller.Service.HistoryPersegi(c)
	utils.PanicIfNeeded(err)
	return c.JSON(utils.ResponseData{
		Code:    200,
		Message: "History persegi",
		Results: response,
	})
}

func (controller CalculationController) HistoryLingkran(c *fiber.Ctx) error {
	response, err := controller.Service.HistoryLingkaran(c)
	utils.PanicIfNeeded(err)
	return c.JSON(utils.ResponseData{
		Code:    200,
		Message: "Histoy lingkaran",
		Results: response,
	})
}

func (controller CalculationController) HistorySegitiga(c *fiber.Ctx) error {
	response, err := controller.Service.HistorySegitiga(c)
	utils.PanicIfNeeded(err)
	return c.JSON(utils.ResponseData{
		Code:    200,
		Message: "History Segitiga",
		Results: response,
	})
}
