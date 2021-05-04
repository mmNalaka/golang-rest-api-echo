package services

import (
	"golang-rest-api-echo/pkg/config"
	"golang-rest-api-echo/pkg/data"
	"golang-rest-api-echo/pkg/domain"
	"golang-rest-api-echo/pkg/models"

	"github.com/go-swagger/go-swagger/fixtures/bugs/1719/pkg/models"
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

func (u UserService) CreateAnAccount(user *domain.User) *models.Error {
}
