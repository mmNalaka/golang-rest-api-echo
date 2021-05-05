package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-rest-api-echo/pkg/config"
	"golang-rest-api-echo/pkg/crypto"
	"golang-rest-api-echo/pkg/data"
	"golang-rest-api-echo/pkg/domain"
	"golang-rest-api-echo/pkg/models"
	"net/http"
)

type IUserService interface {
	Signup(user *domain.User) *models.Error
	Login(user *domain.User) (string, *models.Error)
}

type UserService struct {
	userProvider data.IUserProvider
	cfg          *config.AppConfig
}

func NewUserService(cfg *config.AppConfig, userProvider data.IUserProvider) *UserService {
	return &UserService{
		userProvider: userProvider,
		cfg:          cfg,
	}
}

func (u UserService) CreateAccount(user *domain.User) *models.Error {
	userExist, err := u.userProvider.UsernameExists(user.UserName)

	if err != nil {
		return &models.Error{
			Code:    http.StatusInternalServerError,
			Name:    "SERVER_ERROR",
			Message: "Internal server error!",
			Error:   err,
		}
	}

	if userExist {
		return &models.Error{
			Code:    http.StatusUnprocessableEntity,
			Name:    "USERNAME_TAKEN",
			Message: "This username is already in use.",
		}
	}

	user.ID = primitive.NewObjectID()
	hash, err := crypto.HashPassword(user.Password)

	if err != nil {
		return &models.Error{
			Code:    http.StatusInternalServerError,
			Name:    "SERVER_ERROR",
			Message: "Internal server error!",
			Error:   err,
		}
	}

	user.Password = hash
	err = u.userProvider.CreateAccount(user)

	if err != nil {
		return &models.Error{
			Code:    http.StatusInternalServerError,
			Name:    "SERVER_ERROR",
			Message: "Internal server error!",
			Error:   err,
		}
	}

	return nil
}

func (u UserService) Login(user *domain.User) (string, *models.Error) {
	existingUser, err := u.userProvider.FindByUsername(user.UserName)

	if err != nil {
		return "", &models.Error{
			Code:    http.StatusInternalServerError,
			Name:    "SERVER_ERROR",
			Message: "Internal server error!",
			Error:   err,
		}
	}

	if existingUser == nil {
		return "", &models.Error{
			Code:    http.StatusBadRequest,
			Name:    "INVALID_LOGIN",
			Message: "Invalid username and password",
		}
	}

	err = crypto.ComparePasswordWithHash(existingUser.Password, user.Password)
	if err != nil {
		return "", &models.Error{
			Code:    http.StatusBadRequest,
			Name:    "INVALID_LOGIN",
			Message: "Invalid username and password",
		}
	}

	token, err := crypto.CreateJwtToken(existingUser.ID.Hex(), *u.cfg)
	return token, nil
}
