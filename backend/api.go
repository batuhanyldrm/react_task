package main

import "github.com/gofiber/fiber/v2"

type Api struct {
	Service *Service
}

func NewApi(service *Service) Api {
	return Api{
		Service: service,
	}
}

func (api *Api) GetStockHandler(c *fiber.Ctx) error {

	stock, err := api.Service.GetStock()

	switch err {
	case nil:
		c.JSON(stock)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}
