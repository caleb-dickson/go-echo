package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/security"
)

func IsAuthorized(c echo.Context, page string) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*security.JwtCustomClaims)

	req := c.Request()

	switch req.Method {
	case "GET":
		err := checkPermissions(claims, "view", page)
		if err != nil {
			return err
		}
	case "PUT":
		err := checkPermissions(claims, "edit", page)
		if err != nil {
			return err
		}
	case "PATCH":
		err := checkPermissions(claims, "edit", page)
		if err != nil {
			return err
		}
	case "DELETE":
		err := checkPermissions(claims, "delete", page)
		if err != nil {
			return err
		}
	}

	return nil

	//if req.Method == "GET" {
	//	for _, permission := range claims.Role.Permissions {
	//		if permission.Name == "view_"+page || permission.Name == "edit_"+page {
	//			return nil
	//		}
	//	}
	//} else {
	//	for _, permission := range claims.Role.Permissions {
	//		if permission.Name == "edit_"+page {
	//			return nil
	//		}
	//	}
	//}
	//return errors.New("unauthorized")
}

func checkPermissions(claims *security.JwtCustomClaims, method string, page string) error {
	for _, permission := range claims.Role.Permissions {
		if permission.Name == method+"_"+page {
			return nil
		}
	}
	return errors.New("unauthorized")
}
