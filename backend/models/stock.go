package models

import "time"

type Product struct {
	ID          string    `json:"id" bson:"id"`
	ProductName string    `json:"productName" bson:"productName"`
	Amount      int       `json:"amount" bson:"amount"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
type ProductDTO struct {
	ProductName string    `json:"productName" bson:"productName"`
	Amount      int       `json:"amount" bson:"amount"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
