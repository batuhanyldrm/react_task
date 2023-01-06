package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"example.com/greetings/models"
	"github.com/gofiber/fiber/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetProducts(t *testing.T) {
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
			req, _ := http.NewRequest(http.MethodGet, "/stocks", nil)

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

func TestGetProduct(t *testing.T) {
	Convey("Get stock", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		stock := models.Product{
			ID:          GenerateUUID(8),
			ProductName: "Sweater",
			Amount:      5,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		repository.CreateProduct(stock)

		Convey("When the get request sent", func() {
			app := SetupApp(&api)
			req, _ := http.NewRequest(http.MethodGet, "/stocks/"+stock.ID, nil)
			resp, err := app.Test(req, 3000)

			So(err, ShouldBeNil)

			Convey("Then status code should be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})

			Convey("Then product should be returned", func() {
				actualResult := models.Product{}
				actualRespBody, _ := ioutil.ReadAll(resp.Body)
				err := json.Unmarshal(actualRespBody, &actualResult)

				So(err, ShouldBeNil)

				So(actualResult.ID, ShouldEqual, stock.ID)
				So(actualResult.ProductName, ShouldEqual, stock.ProductName)
				So(actualResult.Amount, ShouldEqual, stock.Amount)
				So(actualResult.CreatedAt, ShouldEqual, stock.CreatedAt)
				So(actualResult.UpdatedAt, ShouldEqual, stock.UpdatedAt)
			})
		})
	})
}

func TestUpdateProduct(t *testing.T) {
	Convey("Update stock", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		stock := models.Product{
			ID:          GenerateUUID(8),
			ProductName: "product",
			Amount:      100,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		repository.CreateProduct(stock)

		Convey("When the put request sent", func() {
			app := SetupApp(&api)

			stock2 := models.ProductDTO{
				ProductName: "newProduct",
				Amount:      80,
				CreatedAt:   time.Now().UTC().Round(time.Second),
			}
			reqBody, err := json.Marshal(stock2)

			So(err, ShouldBeNil)
			req, _ := http.NewRequest(http.MethodPut, "/stock/"+stock.ID, bytes.NewReader(reqBody))
			req.Header.Add("Content-Type", "application/json")

			resp, err := app.Test(req, 3000)

			So(err, ShouldBeNil)

			Convey("then status should be 200", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusOK)
			})

			Convey("Then product stock should be updated", func() {
				actualResult := models.Product{}
				actualRespBody, _ := ioutil.ReadAll(resp.Body)
				err = json.Unmarshal(actualRespBody, &actualResult)
				fmt.Println(actualResult, "bababab")
				So(err, ShouldBeNil)
				So(actualResult.ID, ShouldEqual, stock.ID)
				So(actualResult.ProductName, ShouldEqual, stock2.ProductName)
				So(actualResult.Amount, ShouldEqual, stock2.Amount)
				So(actualResult.CreatedAt, ShouldEqual, stock2.CreatedAt)

			})
		})
	})
}
