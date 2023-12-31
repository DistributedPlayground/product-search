package repository

import (
	"context"

	"github.com/DistributedPlayground/product-search/config"
	"github.com/DistributedPlayground/product-search/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product interface {
	GetById(ctx context.Context, id string) (*model.Product, error)
	List(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]*model.Product, error)
	InsertOne(ctx context.Context, product *model.Product) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter bson.M, update bson.M) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error)
}

type product[T any] struct {
	Base[T]
}

func NewProduct(client *mongo.Client) Product {
	c := client.Database(config.Var.DB_NAME).Collection("product")
	return &product[model.Product]{
		Base: Base[model.Product]{
			Collection: c,
		},
	}
}
