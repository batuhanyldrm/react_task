package main

import (
	"strings"
	"time"

	"example.com/greetings/models"
	"github.com/google/uuid"
)

type Service struct {
	Repository *Repository
}

func NewService(Repository *Repository) Service {
	return Service{
		Repository: Repository,
	}
}

func (service *Service) GetStocks() ([]models.Product, error) {

	stocks, err := service.Repository.GetStocks()

	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (service *Service) GetSearch(query string) ([]models.Product, error) {

	stocks, err := service.Repository.GetProductsWithQuery(query)

	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (service *Service) GetStock(ID string) (models.Product, error) {

	stock, err := service.Repository.GetStock(ID)

	if err != nil {
		return models.Product{}, nil
	}

	return stock, nil
}

func (service *Service) UpdateStocksAmount(productDTO models.ProductDTO, ID string) (*models.Product, error) {

	stock, err := service.Repository.GetStock(ID)
	if err != nil {
		return nil, err
	}

	stock.Amount = stock.Amount - productDTO.Amount
	stock.UpdatedAt = time.Now().UTC().Round(time.Second)

	err = service.Repository.UpdateStocks(stock, ID)

	if err != nil {
		return nil, err
	}

	return &stock, nil
}

func (service *Service) UpdateStocks(productDTO models.ProductDTO, ID string) (*models.Product, error) {

	stock, err := service.Repository.GetStock(ID)
	if err != nil {
		return nil, err
	}

	stock.ProductName = productDTO.ProductName
	stock.Description = productDTO.Description
	stock.Price = productDTO.Price
	stock.Amount = productDTO.Amount
	stock.UpdatedAt = time.Now().UTC().Round(time.Second)

	err = service.Repository.UpdateStocks(stock, ID)

	if err != nil {
		return nil, err
	}

	return &stock, nil
}

func (service *Service) PostStocks(productDTO models.ProductDTO) *models.Product {

	stock := models.Product{}
	stock.ID = GenerateUUID(8)
	stock.CreatedAt = time.Now().UTC().Round(time.Second)
	stock.UpdatedAt = time.Now().UTC().Round(time.Second)
	stock.ProductName = productDTO.ProductName
	stock.Description = productDTO.Description
	stock.Price = productDTO.Price
	stock.Amount = productDTO.Amount

	err := service.Repository.PostStocks(stock)
	if err != nil {
		return nil
	}

	return &stock
}

func (service *Service) DeleteStocks(stockId string) error {

	err := service.Repository.DeleteStocks(stockId)

	if err != nil {
		return err
	}

	return nil
}

func GenerateUUID(length int) string {
	uuid := uuid.New().String()

	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}
	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}
