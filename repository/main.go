package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"pa-back/config"
)

// Clinet mongo Clinet
var db *mongo.Database

func init() {
	ctx := context.TODO()
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		config.JSONConfig.DB.User,
		config.JSONConfig.DB.Password,
		config.JSONConfig.DB.Host,
		config.JSONConfig.DB.Port,
		config.JSONConfig.DB.DBName)
	mongoClinet, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		//mongoClinet.Disconnect(ctx)
		log.Fatal(err)
	}
	ctx = context.TODO()
	err = mongoClinet.Ping(ctx, readpref.Primary())
	if err != nil {
		//mongoClinet.Disconnect(ctx)
		log.Fatal(err)
	}
	db = mongoClinet.Database(config.JSONConfig.DB.DBName)
}
