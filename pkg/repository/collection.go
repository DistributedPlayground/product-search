package repository

import (
	"github.com/DistributedPlayground/product-search/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection interface {
}

type collection[T any] struct {
	Base[T]
}

func NewCollection(c *mongo.Collection) Collection {
	return &collection[model.Collection]{
		Base: Base[model.Collection]{
			Collection: c,
		},
	}
}
