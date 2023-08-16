package service

import (
	"context"

	"github.com/DistributedPlayground/product-search/graph/model"
	"github.com/DistributedPlayground/product-search/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection interface {
	GetById(ctx context.Context, id string) (collection *model.Collection, err error)
	List(ctx context.Context, limit *int, offset *int) (collections []*model.Collection, err error)
}

type collection struct {
	repo repository.Collection
}

func NewCollection(repo repository.Collection) Collection {
	return &collection{
		repo: repo,
	}
}

func (c *collection) GetById(ctx context.Context, id string) (collection *model.Collection, err error) {
	collection, err = c.repo.GetById(ctx, id)
	if err != nil {
		return collection, err
	}
	return collection, nil
}

func (c *collection) List(ctx context.Context, limit *int, offset *int) (collections []*model.Collection, err error) {
	filter := bson.M{}
	opts := options.Find().SetLimit(int64(*limit)).SetSkip(int64(*offset))
	collections, err = c.repo.List(ctx, filter, opts)
	if err != nil {
		return collections, err
	}
	return collections, nil
}
