package controller

import (
	"encoding/csv"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/data/model"
	"go-ng/cmd/api/data/model/entity"
	"net/http"
	"os"
	"strconv"
)

func AllOrders(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))

	return c.JSON(
		http.StatusOK,
		model.Paginate(c, &entity.Order{}, page),
	)
}

func Export(c echo.Context) error {
	filePath := "./cmd/csv/orders.csv"

	if err := createCsv(filePath); err != nil {
		return err
	}

	return c.Attachment(filePath, filePath)
}

type Sales struct {
	Date string `json:"date"`
	Sum  string `json:"sum"`
}

func Chart(c echo.Context) error {
	var sales []Sales

	model.DB.Raw(`
			SELECT to_char(o.created_at, '%Y-%m-%d') as date, sum(oi.price * oi.quantity) as sum
			FROM orders AS o
			JOIN order_items oi on o.id = oi.order_id
			GROUP BY date
		`).Scan(&sales)

	return c.JSON(http.StatusOK, sales)
}

func createCsv(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	var orders []entity.Order

	model.DB.Preload("OrderItems").Find(&orders)

	writer.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})

	for _, order := range orders {
		data := []string{
			strconv.Itoa(int(order.ID)),
			fmt.Sprintf("%v %v", order.FirstName, order.LastName),
			order.Email,
			"",
			"",
			"",
		}

		if err := writer.Write(data); err != nil {
			return err
		}

		for _, item := range order.OrderItems {
			data := []string{
				"",
				"",
				"",
				item.ProductTitle,
				strconv.Itoa(int(item.Price)),
				strconv.Itoa(int(item.Quantity)),
			}

			if err := writer.Write(data); err != nil {
				return err
			}

		}

	}
	return nil
}
