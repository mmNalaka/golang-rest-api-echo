package main

import (
	"golang-rest-api-echo/internal/api"
	"golang-rest-api-echo/pkg/config"
	"golang-rest-api-echo/pkg/db"
)

func main() {
	cfg := config.New()
	db := db.NewMongoConnection(cfg)
	defer db.Disconnect()

	application := api.New()
	application.Start()
}
