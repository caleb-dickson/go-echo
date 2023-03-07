package model

import (
	"github.com/labstack/gommon/log"
	"go-ng/cmd/api/data/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn :=
		"host=localhost user=postgres password=bondstone dbname=gin-gorm " +
			"port=5432 sslmode=disable TimeZone=US/Eastern"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Permission{},
		&entity.Product{},
		&entity.Order{},
		&entity.OrderItem{},
	)
	if err != nil {
		log.Errorf("Failed to AutoMigrate Postgres database: %v", err)
	}

	DB = database
}
