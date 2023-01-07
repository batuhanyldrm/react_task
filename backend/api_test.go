package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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
			Description: "trertte",
			Price:       565,
			Amount:      500,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		stock2 := models.Product{
			ID:          GenerateUUID(8),
			ProductName: "Coat",
			Description: "utjhng",
			Price:       766,
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
				So(actualResult[0].Price, ShouldEqual, stock.Price)
				So(actualResult[0].Description, ShouldEqual, stock.Description)
				So(actualResult[0].Amount, ShouldEqual, stock.Amount)
				So(actualResult[0].CreatedAt, ShouldEqual, stock.CreatedAt)
				So(actualResult[0].UpdatedAt, ShouldEqual, stock.UpdatedAt)
				So(actualResult[1].ID, ShouldEqual, stock2.ID)
				So(actualResult[1].ProductName, ShouldEqual, stock2.ProductName)
				So(actualResult[1].Description, ShouldEqual, stock2.Description)
				So(actualResult[1].Price, ShouldEqual, stock2.Price)
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
			Description: "weqqwqe",
			Price:       455,
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
				So(actualResult.Description, ShouldEqual, stock.Description)
				So(actualResult.Price, ShouldEqual, stock.Price)
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
			Description: "qwqwqw",
			Price:       200,
			Amount:      100,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		repository.CreateProduct(stock)

		Convey("When the put request sent", func() {
			app := SetupApp(&api)

			stock2 := models.ProductDTO{
				ProductName: "newProduct",
				Description: "assasas",
				Price:       232,
				Amount:      80,
				CreatedAt:   time.Now().UTC().Round(time.Second),
			}
			reqBody, err := json.Marshal(stock2)

			So(err, ShouldBeNil)
			req, _ := http.NewRequest(http.MethodPut, "/stocks/"+stock.ID, bytes.NewReader(reqBody))
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
				So(err, ShouldBeNil)
				So(actualResult.ID, ShouldEqual, stock.ID)
				So(actualResult.ProductName, ShouldEqual, stock2.ProductName)
				So(actualResult.Description, ShouldEqual, stock2.Description)
				So(actualResult.Price, ShouldEqual, stock2.Price)
				So(actualResult.Amount, ShouldEqual, stock2.Amount)
				So(actualResult.CreatedAt, ShouldEqual, stock2.CreatedAt)

			})
		})
	})
}

func TestAddProduct(t *testing.T) {
	Convey("Add stock", t, func() {
		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		stock := models.Product{
			ProductName: "Product",
			Description: "sdaasdas",
			Price:       500,
			Amount:      320,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}

		Convey("When the post request sent", func() {

			reqBody, err := json.Marshal(stock)

			req, _ := http.NewRequest(http.MethodPost, "/stocks", bytes.NewReader(reqBody))
			req.Header.Add("Content-Type", "application/json")
			req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))

			app := SetupApp(&api)
			resp, err := app.Test(req, 30000)
			So(err, ShouldBeNil)

			Convey("Then status code should be 201", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusCreated)
			})

			Convey("Then added stock should return", func() {
				actualResult, err := repository.GetStock(stock.ID)

				So(err, ShouldBeNil)
				So(actualResult, ShouldNotBeNil)
				So(actualResult.ProductName, ShouldEqual, stock.ProductName)
				So(actualResult.Description, ShouldEqual, stock.Description)
				So(actualResult.Price, ShouldEqual, stock.Price)
				So(actualResult.Amount, ShouldEqual, stock.Amount)
			})
		})
	})
}
func TestDeleteProduct(t *testing.T) {
	Convey("Delete stock that user wants", t, func() {

		repository := GetCleanTestRepository()
		service := NewService(repository)
		api := NewApi(&service)

		stock := models.Product{
			ID:          GenerateUUID(8),
			ProductName: "Product",
			Description: "kfbgkfjgb",
			Price:       709,
			Amount:      250,
			CreatedAt:   time.Now().UTC().Round(time.Second),
			UpdatedAt:   time.Now().UTC().Round(time.Second),
		}
		repository.CreateProduct(stock)

		Convey("When the delete request sent", func() {
			app := SetupApp(&api)

			req, _ := http.NewRequest(http.MethodDelete, "/stocks/"+stock.ID, nil)
			resp, err := app.Test(req, 30000)
			So(err, ShouldBeNil)

			Convey("Then status code should be 204", func() {
				So(resp.StatusCode, ShouldEqual, fiber.StatusNoContent)
			})

			Convey("Then stock should be deleted", func() {
				stock, _ := repository.GetStocks()
				So(stock, ShouldHaveLength, 0)
				So(stock, ShouldResemble, []models.Product{})
			})
		})
	})
}
