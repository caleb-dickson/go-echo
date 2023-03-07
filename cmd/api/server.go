package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-ng/cmd/api/data/bootstrap"
	"go-ng/cmd/api/data/model"
	"go-ng/cmd/api/data/model/validation"
	"go-ng/cmd/api/logging"
	"go-ng/cmd/api/route"
)

func main() {
	// create Echo app
	e := echo.New()

	// assign custom validator
	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	// config custom logging setup
	logging.CustomLoggerConfig(e)

	e.Use(middleware.Recover())

	// CORS config
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	// connect to database
	model.ConnectDatabase()

	// register routes
	route.Setup(e)

	// seeder
	bootstrap.LoadPermissions()

	e.Logger.Fatal(e.Start(":1323"))
}
