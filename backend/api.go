package main

import (
	"example.com/greetings/models"
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

	stocks, err := api.Service.GetStocks()

	switch err {
	case nil:
		c.JSON(stocks)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) GetSearchHandler(c *fiber.Ctx) error {

	query := c.Query("q")

	stocks, err := api.Service.GetSearch(query)

	switch err {
	case nil:
		c.JSON(stocks)
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

func (api *Api) UpdateStocksAmountHandler(c *fiber.Ctx) error {

	ID := c.Params("id")
	stock := models.ProductDTO{}
	err := c.BodyParser(&stock)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	updatedStock, err := api.Service.UpdateStocksAmount(stock, ID)

	switch err {
	case nil:
		c.JSON(updatedStock)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) UpdateStocksHandler(c *fiber.Ctx) error {

	ID := c.Params("id")
	stock := models.ProductDTO{}
	err := c.BodyParser(&stock)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}

	updatedStock, err := api.Service.UpdateStocks(stock, ID)

	switch err {
	case nil:
		c.JSON(updatedStock)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) PostStocksHandler(c *fiber.Ctx) error {

	createStocks := models.ProductDTO{}
	err := c.BodyParser(&createStocks)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
	}
	stock := api.Service.PostStocks(createStocks)

	switch err {
	case nil:
		c.JSON(stock)
		c.Status(fiber.StatusCreated)

	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (api *Api) DeleteStocksHandler(c *fiber.Ctx) error {

	ID := c.Params("id")
	err := api.Service.DeleteStocks(ID)

	switch err {
	case nil:
		c.Status(fiber.StatusNoContent)

	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}
