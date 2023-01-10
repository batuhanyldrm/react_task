package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	repository := NewRepository()
	service := NewService(repository)
	api := NewApi(&service)
	app := SetupApp(&api)
	app.Listen(":3001")
}

func SetupApp(api *Api) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/stocks", api.GetStocksHandler)
	app.Put("stocks/:id/amount", api.UpdateStocksAmountHandler)
	app.Get("/stocks/:id", api.GetStockHandler)
	app.Put("/stocks/:id", api.UpdateStocksHandler)
	app.Post("/stocks", api.PostStocksHandler)
	app.Delete("/stocks/:id", api.DeleteStocksHandler)
	app.Get("/search", api.GetSearchHandler)

	return app
}
