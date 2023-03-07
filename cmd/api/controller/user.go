package controller

import (
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/data/model"
	"go-ng/cmd/api/data/model/entity"
	"go-ng/cmd/api/middleware"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func AllUsers(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))

	return c.JSON(
		http.StatusOK,
		model.Paginate(c, &entity.User{}, page),
	)
}

func CreateUser(c echo.Context) error {
	var reqBodyInput entity.User
	if err := (&echo.DefaultBinder{}).BindBody(c, &reqBodyInput); err != nil {
		return c.String(http.StatusBadRequest, "All fields are required")
	}
	if err := c.Validate(reqBodyInput); err != nil {
		return err
	}

	user := entity.User{
		FirstName: reqBodyInput.FirstName,
		LastName:  reqBodyInput.LastName,
		Email:     reqBodyInput.Email,
		Username:  reqBodyInput.Username,
		RoleId:    1,
	}
	user.SetPasswordHashed("1234")
	model.DB.Create(&user)

	user.RedactPassword()
	return c.JSON(http.StatusAccepted, user)
}

func GetUser(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	user := entity.User{
		Model: gorm.Model{ID: uint(id)},
	}
	model.DB.Preload("Role").Find(&user)

	user.RedactPassword()
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {

	err := middleware.IsAuthorized(c, "users")
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	id, _ := strconv.Atoi(c.Param("id"))

	var reqBodyInput entity.User
	if err := (&echo.DefaultBinder{}).BindBody(c, &reqBodyInput); err != nil {
		return c.String(http.StatusBadRequest, "All fields are required")
	}
	reqBodyInput.ID = uint(id)

	model.DB.Model(&reqBodyInput).Updates(reqBodyInput)

	reqBodyInput.RedactPassword()
	return c.JSON(http.StatusAccepted, reqBodyInput)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := entity.User{
		Model: gorm.Model{ID: uint(id)},
	}

	model.DB.Delete(&user)

	return nil
}
