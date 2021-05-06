package api

import (
	"github.com/labstack/echo/v4"
	"golang-rest-api-echo/pkg/models"
	"net/http"
)

func (a App) Signup(c echo.Context) error {
	newUser, err := models.ValidateSignupRequest(c)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	err = a.userSvc.Signup(newUser)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.String(http.StatusCreated, "")
}

func (a App) Login(c echo.Context) error {
	loginRequest, err := models.ValidateLoginRequest(c)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	token, err := a.userSvc.Login(loginRequest)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	response := &models.LoginResponse{
		Token: token,
	}

	return c.JSON(http.StatusOK, response)
}
