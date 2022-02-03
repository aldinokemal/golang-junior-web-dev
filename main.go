package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"lsp/controllers"
	"lsp/repositories/csv"
	"lsp/services"
)

func main() {
	// Setup Repository
	persegiRepository := csv.NewPersegiRepository()
	lingkaranRepository := csv.NewLingkaranRepository()
	segitigaRepository := csv.NewSegitigaRepository()

	// Setup Services
	calculateService := services.NewCalculationService(&persegiRepository, &segitigaRepository, &lingkaranRepository)

	// Setup Controller
	calculateController := controllers.NewCalculationController(calculateService)

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(recover.New())
	calculateController.Route(app)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Running Well 👋!")
	})
	app.Get("/index", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	app.Get("/hitung/persegi", func(c *fiber.Ctx) error {
		return c.Render("calculate_persegi", fiber.Map{})
	})
	app.Get("/hitung/segitiga", func(c *fiber.Ctx) error {
		return c.Render("calculate_segitiga", fiber.Map{})
	})
	app.Get("/hitung/lingkaran", func(c *fiber.Ctx) error {
		return c.Render("calculate_lingkaran", fiber.Map{})
	})
	app.Listen(":3000")
}
