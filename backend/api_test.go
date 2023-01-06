package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"example.com/greetings/models"
	"github.com/gofiber/fiber/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetProduct(t *testing.T) {
	Convey("Get Stocks", t, func() {
		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		stock := models.Product{
			ID:          GenerateUUID(8),
			ProductName: "Pants",
			Amount:      500,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		stock2 := models.Product{
			ID:          GenerateUUID(8),
			ProductName: "Coat",
			Amount:      499,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		repository.CreateProduct(stock)
		repository.CreateProduct(stock2)

		Convey("When the get reguest sent", func() {
			app := SetupApp(&api)
			req, _ := http.NewRequest(http.MethodGet, "/stock", nil)

			resp, err := app.Test(req)
			So(err, ShouldBeNil)

			Convey("Then status code shuld be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})

			Convey("Then the request should return all products", func() {
				actualResult := []models.Product{}
				actualResponseBody, _ := ioutil.ReadAll(resp.Body)
				err := json.Unmarshal(actualResponseBody, &actualResult)

				So(err, ShouldBeNil)

				So(actualResult, ShouldHaveLength, 2)
				So(actualResult[0].ID, ShouldEqual, stock.ID)
				So(actualResult[0].ProductName, ShouldEqual, stock.ProductName)
				So(actualResult[0].Amount, ShouldEqual, stock.Amount)
				So(actualResult[0].CreatedAt, ShouldEqual, stock.CreatedAt)
				So(actualResult[0].UpdatedAt, ShouldEqual, stock.UpdatedAt)
				So(actualResult[1].ID, ShouldEqual, stock2.ID)
				So(actualResult[1].ProductName, ShouldEqual, stock2.ProductName)
				So(actualResult[1].Amount, ShouldEqual, stock2.Amount)
				So(actualResult[1].CreatedAt, ShouldEqual, stock2.CreatedAt)
				So(actualResult[1].UpdatedAt, ShouldEqual, stock2.UpdatedAt)
			})
		})
	})
}
