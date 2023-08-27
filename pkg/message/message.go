package message

import (
	"github.com/DistributedPlayground/product-search/pkg/repository"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Messages struct {
	Collection Collection
	Product    Product
}

func NewMessages(kc *kafka.Consumer, repos repository.Repositories) Messages {
	return Messages{
		Collection: NewCollection(kc, repos.Collection),
		Product:    NewProduct(kc, repos.Product),
	}
}
