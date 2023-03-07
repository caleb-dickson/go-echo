package security

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"go-ng/cmd/api/data/model/entity"
)

const (
	// AuthCookieName , found in api/security/main.go is the key name of the cookie that stores the JWT
	AuthCookieName = "sP_7I"
	Key            = "Go is an open source programming language supported by Google"
)

var JwtConfig = middleware.JWTConfig{
	TokenLookup: fmt.Sprintf("cookie:%v", AuthCookieName),
	Claims:      &JwtCustomClaims{},
	SigningKey:  []byte(Key),
}

type JwtCustomClaims struct {
	ID        uint        `json:"id"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	Email     string      `json:"email"`
	Username  string      `json:"username"`
	Role      entity.Role `json:"role"`
	jwt.StandardClaims
}

type Claims struct {
	jwt.StandardClaims
}
