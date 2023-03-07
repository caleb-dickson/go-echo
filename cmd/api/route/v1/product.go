package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-ng/cmd/api/controller"
	"go-ng/cmd/api/security"
)

func SetupProductRoutes(public *echo.Group) {

	productsBase := public.Group("/products")

	productsBase.GET("", controller.AllProducts)
	productsBase.GET("/:id", controller.GetProduct)

	secured := public.Group("/s/products")
	secured.Use(middleware.JWTWithConfig(security.JwtConfig))

	secured.POST("", controller.CreateProduct)
	secured.PUT("/:id", controller.UpdateProduct)
	secured.DELETE("/:id", controller.DeleteProduct)

}
