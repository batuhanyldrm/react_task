package main

import (
	"context"
	"log"
	"time"

	"example.com/greetings/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func (repository *Repository) CreateProduct(product models.Product) error {
	collection := repository.client.Database("stock").Collection("stock")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, product)

	if err != nil {
		return err
	}

	return nil
}

func NewRepository() *Repository {
	uri := "mongodb+srv://Cluster:bthn998877@cluster0.hnmuy.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}

func NewTestRepository() *Repository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}

func (repository *Repository) GetStock() ([]models.Product, error) {
	collection := repository.client.Database("stock").Collection("stock")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	stocks := []models.Product{}
	for cur.Next(ctx) {
		var stock models.Product
		err := cur.Decode(&stock)
		if err != nil {
			log.Fatal(err)
		}

		stocks = append(stocks, stock)
	}

	return stocks, nil

}

func GetCleanTestRepository() *Repository {

	repository := NewRepository()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	stockDB := repository.client.Database("stock")
	stockDB.Drop(ctx)

	return repository
}
