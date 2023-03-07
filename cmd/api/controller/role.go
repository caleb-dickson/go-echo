package controller

import (
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/data/model"
	"go-ng/cmd/api/data/model/entity"
	"go-ng/cmd/api/data/model/request"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func AllRoles(c echo.Context) error {
	var roles []entity.Role

	model.DB.Preload("Permissions").Find(&roles)

	return c.JSON(http.StatusOK, roles)
}

func CreateRole(c echo.Context) error {
	var reqBodyInput request.RoleRequestDto
	if err := (&echo.DefaultBinder{}).BindBody(c, &reqBodyInput); err != nil {
		return c.String(http.StatusBadRequest, "All fields are required")
	}
	if err := c.Validate(reqBodyInput); err != nil {
		return err
	}

	permissions := make([]entity.Permission, len(reqBodyInput.Permissions))

	for i, permId := range reqBodyInput.Permissions {
		id, _ := strconv.Atoi(permId)
		permissions[i] = entity.Permission{
			Model: gorm.Model{ID: uint(id)},
		}
	}

	model.DB.Where("id IN ?", reqBodyInput.Permissions).Find(&permissions)

	role := entity.Role{
		Name:        reqBodyInput.Name,
		Permissions: permissions,
	}
	model.DB.Preload("Permissions").Create(&role)

	return c.JSON(http.StatusAccepted, role)
}

func GetRole(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	role := entity.Role{
		Model: gorm.Model{ID: uint(id)},
	}
	model.DB.Preload("Permissions").Find(&role)

	return c.JSON(http.StatusOK, role)
}

func UpdateRole(c echo.Context) error {
	paramId, _ := strconv.Atoi(c.Param("id"))

	var reqBodyInput request.RoleRequestDto
	if err := (&echo.DefaultBinder{}).BindBody(c, &reqBodyInput); err != nil {
		return c.String(http.StatusBadRequest, "All fields are required")
	}

	permissions := make([]entity.Permission, len(reqBodyInput.Permissions))

	for i, permId := range reqBodyInput.Permissions {
		id, _ := strconv.Atoi(permId)
		permissions[i] = entity.Permission{
			Model: gorm.Model{ID: uint(id)},
		}
	}

	perms := model.DB.Where("id IN ?", reqBodyInput.Permissions).Find(&permissions)
	foundPerms := perms.RowsAffected

	if len(reqBodyInput.Permissions) > 0 && foundPerms > 0 {
		var res interface{}
		model.DB.Table("roles_permissions").Where("role_id", paramId).Delete(res)
	}

	role := entity.Role{
		Model:       gorm.Model{ID: uint(paramId)},
		Name:        reqBodyInput.Name,
		Permissions: permissions,
	}

	model.DB.Model(&role).Updates(role)

	return c.JSON(http.StatusAccepted, role)
}

func DeleteRole(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	role := entity.Role{
		Model: gorm.Model{ID: uint(id)},
	}

	model.DB.Delete(&role)
	return nil
}
