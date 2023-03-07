package v1

import (
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/controller"
)

func SetupImageRoutes(imagesRoute *echo.Group) {

	// secured routes
	imagesRoute.POST("/upload", controller.Upload)
	imagesRoute.Static("/uploads", "./cmd/uploads") // Path Prefix ex. = "/uploads/image.png"
}
