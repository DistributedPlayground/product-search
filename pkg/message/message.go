package message

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type Messages struct {
	Collection Collection
	Product    Product
}

func NewMessages(kc *kafka.Consumer) Messages {
	return Messages{
		Collection: NewCollection(kc),
		Product:    NewProduct(kc),
	}
}
