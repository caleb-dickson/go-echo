package v1

import (
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/controller"
)

func SetupPermissionRoutes(permsRoute *echo.Group) {
	permsRoute.GET("", controller.AllPermissions)
}
