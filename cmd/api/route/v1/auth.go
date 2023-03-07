package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-ng/cmd/api/controller"
	"go-ng/cmd/api/security"
)

func SetupAuthRoutes(public *echo.Group) { // public auth routes
	public.POST("/signup", controller.Signup)
	public.POST("/login", controller.Login)

	// secured auth routes config
	secured := public.Group("/s")
	secured.Use(middleware.JWTWithConfig(security.JwtConfig))

	// secured routes
	secured.PUT("/credentials/:id", controller.UpdateCredentials)
	secured.GET("", controller.TestAuth)
	secured.PUT("/logout", controller.Logout)
}
