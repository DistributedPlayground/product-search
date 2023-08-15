package store

import (
	"context"
	"fmt"
	"time"

	"github.com/DistributedPlayground/products/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func mongoConnectionString() string {
	var (
		DBUser     = config.Var.DB_USERNAME
		DBPassword = config.Var.DB_PASSWORD
		DBName     = config.Var.DB_NAME
		DBHost     = config.Var.DB_HOST
		DBPort     = config.Var.DB_PORT
	)

	str := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName,
	)

	return str
}

func MustNewMongo() *mongo.Client {
	if mongoClient != nil {
		return mongoClient
	}

	clientOptions := options.Client().ApplyURI(mongoConnectionString())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	mongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return mongoClient
}
