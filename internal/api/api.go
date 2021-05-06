package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-rest-api-echo/pkg/config"
	"golang-rest-api-echo/pkg/data"
	"golang-rest-api-echo/pkg/services"
	"net/http"
)

type App struct {
	server  *echo.Echo
	userSvc services.IUserService
	cgf     *config.AppConfig
}

func New(cfg *config.AppConfig, client *mongo.Client) *App {
	server := echo.New()

	// providers
	userProvider := data.NewUserProvider(cfg, client)

	// services
	userSvc := services.NewUserService(cfg, userProvider)

	return &App{
		server:  server,
		userSvc: userSvc,
		cgf:     cfg,
	}
}

func (a App) RegisterMiddlewares() {
	a.server.Use(middleware.Logger())
	a.server.Use(middleware.Recover())
	a.server.Use(middleware.RequestID())
}

func (a App) RegisterRouters() {
	a.server.GET("/api/v1/health", a.HeathCheck)
	a.server.POST("/api/v1/auth/signup", a.Signup)
	a.server.POST("/api/v1/auth/login", a.Login)

	protected := a.server.Group("/api/v1/private")
	middlewares := MiddleWares{config: a.cgf}

	protected.Use(middlewares.Authenticated)
	protected.GET("/secret", func(c echo.Context) error {
		userId := c.Get("user").(string)
		return c.String(http.StatusOK, userId)
	})
}

func (a App) Start() {
	a.RegisterMiddlewares()
	a.RegisterRouters()

	a.server.Start(":5000")
}
