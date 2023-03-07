package controller

import (
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/data/model"
	"go-ng/cmd/api/data/model/entity"
	"net/http"
)

func AllPermissions(c echo.Context) error {
	var permissions []entity.Permission

	model.DB.Preload("Roles").Find(&permissions)

	return c.JSON(http.StatusOK, permissions)
}
