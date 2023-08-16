package graph

import "github.com/DistributedPlayground/product-search/pkg/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	services service.Services
}

func NewResolver(services service.Services) *Resolver {
	return &Resolver{
		services: services,
	}
}
