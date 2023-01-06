package main

import (
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	Service *Service
}

func NewApi(service *Service) Api {
	return Api{
		Service: service,
	}
}

func (api *Api) GetStocksHandler(c *fiber.Ctx) error {

	Stocks, err := api.Service.GetStocks()

	switch err {
	case nil:
		c.JSON(Stocks)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) GetStockHandler(c *fiber.Ctx) error {
	ID := c.Params("id")
	stock, err := api.Service.GetStock(ID)

	switch err {
	case nil:
		c.JSON(stock)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil

}
