package v1

import (
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/controller"
)

func SetupOderRoutes(ordersRoute *echo.Group) {
	// secured routes
	ordersRoute.GET("", controller.AllOrders)
	ordersRoute.GET("/export", controller.Export)
	ordersRoute.GET("/chart", controller.Chart)
}
