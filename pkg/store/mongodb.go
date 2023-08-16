package store

import (
	"context"
	"fmt"
	"time"

	"github.com/DistributedPlayground/product-search/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func mongoConnectionString() string {
	str := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.Var.DB_USERNAME,
		config.Var.DB_PASSWORD,
		config.Var.DB_HOST,
		config.Var.DB_PORT,
	)
	fmt.Println(str)

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
