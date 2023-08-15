package repository

import (
	"github.com/DistributedPlayground/product-search/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product interface {
}

type product[T any] struct {
	Base[T]
}

func NewProduct(c *mongo.Collection) Product {
	return &product[model.Product]{
		Base: Base[model.Product]{
			Collection: c,
		},
	}
}
