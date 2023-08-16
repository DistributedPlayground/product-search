package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repositories struct {
	Collection Collection
	Product    Product
}

func NewRepos(db *mongo.Client) Repositories {
	return Repositories{
		Collection: NewCollection(db),
		Product:    NewProduct(db),
	}
}
