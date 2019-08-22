package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"pa-back/config"
)

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.JSONConfig.DB.User, config.JSONConfig.DB.Password, config.JSONConfig.DB.Host, config.JSONConfig.DB.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	client.Disconnect(ctx)
}
