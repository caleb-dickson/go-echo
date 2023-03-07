package v1

import (
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/controller"
)

func SetupUserRoutes(usersRoute *echo.Group) {
	// secured routes
	usersRoute.GET("", controller.AllUsers)
	usersRoute.POST("", controller.CreateUser)
	usersRoute.GET("/:id", controller.GetUser)
	usersRoute.PUT("/:id", controller.UpdateUser)
	usersRoute.DELETE("/:id", controller.DeleteUser)
}
