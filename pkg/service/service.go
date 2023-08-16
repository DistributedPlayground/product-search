package service

import (
	"github.com/DistributedPlayground/product-search/pkg/repository"
)

type Services struct {
	Collection Collection
	Product    Product
}

func NewServices(repos repository.Repositories) Services {
	return Services{
		Collection: NewCollection(repos.Collection),
		Product:    NewProduct(repos.Product),
	}
}
