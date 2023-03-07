package logging

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CustomLoggerConfig(e *echo.Echo) {

	count := 0

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		BeforeNextFunc: func(c echo.Context) {
			count++
			c.Set("customValueFromContext", count)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST# %v: %v %v, status %v, %v\n", value, v.Method, v.URI, v.Status, v.Latency)
			return nil
		},
		LogLatency: true,
		LogMethod:  true,
		LogURI:     true,
		LogStatus:  true,
	}))
}
