package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/data/model"
	"go-ng/cmd/api/data/model/entity"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func AllProducts(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))

	return c.JSON(
		http.StatusOK,
		model.Paginate(c, &entity.Product{}, page),
	)
}

func CreateProduct(c echo.Context) error {
	var reqBodyInput entity.Product
	if err := (&echo.DefaultBinder{}).BindBody(c, &reqBodyInput); err != nil {
		return c.String(http.StatusBadRequest, "Missing required fields")
	}
	if err := c.Validate(reqBodyInput); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Missing required fields")
	}

	model.DB.Create(&reqBodyInput)

	return c.JSON(http.StatusAccepted, reqBodyInput)
}

func GetProduct(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	product := entity.Product{
		Model: gorm.Model{ID: uint(id)},
	}
	model.DB.Find(&product)

	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var reqBodyInput entity.Product
	if err := (&echo.DefaultBinder{}).BindBody(c, &reqBodyInput); err != nil {
		return c.String(http.StatusBadRequest, "All fields are required")
	}
	reqBodyInput.ID = uint(id)

	model.DB.Model(&reqBodyInput).Updates(reqBodyInput)

	return c.JSON(http.StatusAccepted, reqBodyInput)
}

func DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product := entity.Product{
		Model: gorm.Model{ID: uint(id)},
	}

	model.DB.Delete(&product)
	return nil
}
