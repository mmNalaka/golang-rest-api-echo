package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang-rest-api-echo/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	Client *mongo.Client
	ctx    context.Context
}

func NewMongoConnection(cfg *config.AppConfig) Connection {
	uri := fmt.Sprintf("mongodb://%s/%s", cfg.DbHost, cfg.DbName)

	credentialed := options.Credential{
		Username: cfg.DbUser,
		Password: cfg.DbPassword,
	}

	clientOptions := options.Client().ApplyURI(uri).SetAuth(credentialed)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("MongoDB connection error, ", err)
	}

	fmt.Println("INFO: Connected to database.")

	return Connection{
		Client: client,
		ctx:    ctx,
	}
}

func (c Connection) Disconnect() {
	c.Client.Disconnect(c.ctx)
}
