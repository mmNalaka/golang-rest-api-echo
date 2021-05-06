package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang-rest-api-echo/pkg/config"
)

type MiddleWares struct {
	config *config.AppConfig
}

func (m MiddleWares) Authenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenFromHeader := c.Request().Header.Get("Authorization")
		if tokenFromHeader == "" {
			return echo.ErrUnauthorized
		}

		type Claims struct {
			Id  string `json:"id"`
			Exp int    `json:"exp"`
			jwt.StandardClaims
		}

		token, err := jwt.ParseWithClaims(tokenFromHeader, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.config.JWtSecret), nil
		})

		if err != nil {
			fmt.Println(err)
			return echo.ErrUnauthorized
		}

		claims, ok := token.Claims.(*Claims)
		if ok && token.Valid {
			c.Set("user", claims.Id)
			return next(c)
		} else {
			return echo.ErrUnauthorized
		}
	}
}
