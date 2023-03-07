package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string  `json:"title" validation:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validation:"required"`
}

func (p *Product) Count(db *gorm.DB) int64 {
	var totalProducts int64
	db.Model(&Product{}).Count(&totalProducts)

	return totalProducts
}

func (p *Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []Product
	db.Limit(limit).Offset(offset).Find(&products)

	return products
}
