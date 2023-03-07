package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-ng/cmd/api/route/v1"
	"go-ng/cmd/api/security"
)

func Setup(e *echo.Echo) {
	// API v1
	// public routes group
	public := e.Group("/api/v1")

	// secured routes group
	secured := public.Group("/s")
	secured.Use(middleware.JWTWithConfig(security.JwtConfig))

	// register routes
	v1.SetupAuthRoutes(public)
	v1.SetupUserRoutes(secured.Group("/users"))
	v1.SetupRoleRoutes(secured.Group("/roles"))
	v1.SetupPermissionRoutes(secured.Group("/permissions"))

	v1.SetupProductRoutes(public)
	v1.SetupImageRoutes(secured.Group("/images"))
	v1.SetupOderRoutes(secured.Group("/orders"))
	// API v2??
}
