package store

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var kc *kafka.Consumer

func MustNewKafka() *kafka.Consumer {
	if kc != nil {
		return kc
	}
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "product-search",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	kc = consumer
	return kc
}
