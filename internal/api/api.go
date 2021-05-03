package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	server *echo.Echo
}

func New() *App {
	server := echo.New()
	return &App{
		server: server,
	}
}

func (a App) RegisterMiddlewares() {
	a.server.Use(middleware.Logger())
	a.server.Use(middleware.Recover())
}

func (a App) RegisterRouters() {
	a.server.GET("/api/v1/health", a.HeathCheck)
}

func (a App) Start() {
	a.RegisterMiddlewares()
	a.RegisterRouters()

	a.server.Start(":5000")
}
