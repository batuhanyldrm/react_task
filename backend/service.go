package main

import (
	"strings"

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

func (service *Service) GetStock(ID string) (models.Product, error) {

	stock, err := service.Repository.GetStock(ID)

	if err != nil {
		return models.Product{}, nil
	}

	return stock, nil
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
