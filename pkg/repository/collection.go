package repository

import (
	"context"

	"github.com/DistributedPlayground/product-search/config"
	"github.com/DistributedPlayground/product-search/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection interface {
	GetById(ctx context.Context, id string) (*model.Collection, error)
	List(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]*model.Collection, error)
	InsertOne(ctx context.Context, collection *model.Collection) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter bson.M, update bson.M) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error)
}

type collection[T any] struct {
	Base[T]
}

func NewCollection(client *mongo.Client) Collection {
	c := client.Database(config.Var.DB_NAME).Collection("collection")
	return &collection[model.Collection]{
		Base: Base[model.Collection]{
			Collection: c,
		},
	}
}
