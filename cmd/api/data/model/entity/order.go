package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"`
	Email      string      `json:"email"`
	Total      float32     `json:"total" gorm:"-"`
	OrderItems []OrderItem `json:"orderItems" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	gorm.Model
	OrderId      uint    `json:"orderId"`
	ProductTitle string  `json:"productTitle"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (order *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(order).Count(&total)

	return total
}

func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order

	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	for i, _ := range orders {
		var total float32 = 0

		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float32(orderItem.Quantity)
		}
		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName
		orders[i].Total = total
	}

	return orders
}
