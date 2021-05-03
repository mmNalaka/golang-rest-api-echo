package models

import (
	"golang-rest-api-echo/pkg/domain"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func ValidateSignupRequest(c echo.Context) (*domain.User, *Error) {
	signupRequest := new(SignupRequest)
	if err := c.Bind(signupRequest); err != nil {
		return nil, BindError()
	}

	var validationErrors []string

	if len(signupRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 charactors long")
	}
	if len(signupRequest.Username) < 3 {
		validationErrors = append(validationErrors, "Username must be 3 charactors long")
	}
	if len(signupRequest.Name) < 2 {
		validationErrors = append(validationErrors, "Password must be 2 charactors long")
	}
	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)
	}

	return &domain.User{
		UserName: signupRequest.Username,
		Name:     signupRequest.Name,
		Password: signupRequest.Password,
	}, nil
}

func ValidateLoginRequest(c echo.Context) (*domain.User, *Error) {
	loginRequest := new(LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return nil, BindError()
	}

	var validationErrors []string

	if len(loginRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 charactors long")
	}
	if len(loginRequest.Username) < 3 {
		validationErrors = append(validationErrors, "Username must be 3 charactors long")
	}
	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)
	}

	return &domain.User{
		UserName: loginRequest.Username,
		Password: loginRequest.Password,
	}, nil
}
