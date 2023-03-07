package controller

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-ng/cmd/api/data/model"
	"go-ng/cmd/api/data/model/entity"
	"go-ng/cmd/api/data/model/request"
	"go-ng/cmd/api/data/model/response"
	"go-ng/cmd/api/middleware"
	"go-ng/cmd/api/security"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func Signup(c echo.Context) error {
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
		RoleId:    reqBodyInput.RoleId,
	}
	user.SetPasswordHashed(reqBodyInput.Password)
	model.DB.Create(&user)

	user.RedactPassword()
	user.InsertRolePlaceholderText()
	return c.JSON(http.StatusAccepted, user)
}

func Login(c echo.Context) error {
	var input request.UserCredentialsRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &input); err != nil {
		return c.String(http.StatusBadRequest, "All fields are required")
	}

	var user entity.User
	model.DB.Where("email = ?", input.Email).Preload("Role.Permissions").First(&user)
	if user.ID == 0 {
		return c.JSON(
			http.StatusNotFound,
			response.MessageResponse{
				Message: fmt.Sprintf("User with email '%v' not found", input.Email),
			})
	}

	err := user.AuthenticatePassword(input.Password)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.MessageResponse{
				Message: "Password incorrect",
			})
	}

	// Set custom claims
	claims := &security.JwtCustomClaims{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Role:      user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    strconv.Itoa(int(user.ID)),
			Subject:   user.Username,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tkn, err := token.SignedString([]byte(security.Key))
	if err != nil {
		return err
	}

	// define auth cookie
	cookie := http.Cookie{
		Name:     security.AuthCookieName,
		Value:    tkn,
		Path:     "/api/v1/s",
		Expires:  time.Now().Add(time.Hour * 6),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	// set auth cookie
	c.SetCookie(&cookie)

	user.RedactPassword()
	return c.JSON(http.StatusOK, user)
}

// TestAuth is simply a controller, positioned behind the auth middleware,
// that returns a greeting with the authenticated User's username as proof of authentication
func TestAuth(c echo.Context) error {

	err := middleware.IsAuthorized(c, "users")
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*security.JwtCustomClaims)
	username := claims.Username
	return c.JSON(http.StatusOK, echo.Map{"message": "Welcome " + username + "!"})
}

func Logout(c echo.Context) error { // retrieve the authenticated user
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*security.JwtCustomClaims)
	name := claims.Username

	// invalidate the auth cookie
	cookie := http.Cookie{
		Name:     security.AuthCookieName,
		Value:    "",
		Expires:  time.Now().Add(time.Hour - 1),
		HttpOnly: true,
	}
	c.SetCookie(&cookie)

	// return logout verification
	return c.JSON(
		http.StatusOK,
		response.MessageResponse{
			Message: fmt.Sprintf("'%v' has been logged out", name),
		})
}

func UpdateCredentials(c echo.Context) error {
	paramID, _ := strconv.Atoi(c.Param("id"))

	var reqBodyInput request.UserCredentialsRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &reqBodyInput); err != nil {
		return c.String(http.StatusBadRequest, "All fields are required")
	}

	user := entity.User{Model: gorm.Model{ID: uint(paramID)}}
	model.DB.Preload("Role.Permissions").Find(&user)

	userUpdated := entity.User{
		Model: gorm.Model{ID: user.ID},
		Email: reqBodyInput.Email,
	}
	userUpdated.SetPasswordHashed(reqBodyInput.Password)

	model.DB.Model(&user).Updates(userUpdated)

	user.RedactPassword()
	return c.JSON(http.StatusAccepted, user)
}
