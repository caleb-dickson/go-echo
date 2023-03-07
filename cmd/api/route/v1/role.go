package v1

import (
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/controller"
)

func SetupRoleRoutes(rolesRoute *echo.Group) {
	// secured routes
	rolesRoute.GET("", controller.AllRoles)
	rolesRoute.POST("", controller.CreateRole)
	rolesRoute.GET("/:id", controller.GetRole)
	rolesRoute.PUT("/:id", controller.UpdateRole)
	rolesRoute.DELETE("/:id", controller.DeleteRole)
}
