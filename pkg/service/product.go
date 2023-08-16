package service

import (
	"context"

	"github.com/DistributedPlayground/product-search/graph/model"
	"github.com/DistributedPlayground/product-search/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product interface {
	GetById(ctx context.Context, id string) (product *model.Product, err error)
	List(ctx context.Context, limit *int, offset *int) (products []*model.Product, err error)
}

type product struct {
	repo repository.Product
}

func NewProduct(repo repository.Product) Product {
	return &product{
		repo: repo,
	}
}

func (p *product) GetById(ctx context.Context, id string) (product *model.Product, err error) {
	product, err = p.repo.GetById(ctx, id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *product) List(ctx context.Context, limit *int, offset *int) (products []*model.Product, err error) {
	filter := bson.M{}
	opts := options.Find().SetLimit(int64(*limit)).SetSkip(int64(*offset))
	products, err = p.repo.List(ctx, filter, opts)
	if err != nil {
		return products, err
	}
	return products, nil
}
