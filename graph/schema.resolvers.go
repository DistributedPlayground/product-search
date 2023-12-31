package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	graph "github.com/DistributedPlayground/product-search/graph/api"
	"github.com/DistributedPlayground/product-search/graph/model"
)

// Collections is the resolver for the collections field.
func (r *queryResolver) Collections(ctx context.Context, limit *int, offset *int) ([]*model.Collection, error) {
	collections, err := r.services.Collection.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, limit *int, offset *int) ([]*model.Product, error) {
	products, err := r.services.Product.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Collection is the resolver for the collection field.
func (r *queryResolver) Collection(ctx context.Context, id string) (*model.Collection, error) {
	collection, err := r.services.Collection.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	product, err := r.services.Product.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
