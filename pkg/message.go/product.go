package message

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Product interface {
	Listen() error
}

type product struct {
	kc    *kafka.Consumer
	topic string
}

func NewProduct(kc *kafka.Consumer) Product {
	return &product{kc: kc, topic: "product"}
}
func (p product) Listen() error {
	// subscribe to a topic
	err := p.kc.Subscribe(p.topic, nil)
	if err != nil {
		panic("Failed to subscribe to topic: " + err.Error())
		os.Exit(1)
	}

	// continuously poll for new messages
	for {
		msg, err := p.kc.ReadMessage(-1)
		if err != nil {
			// handle error
			fmt.Printf("Error reading message: %s\n", err)
			continue
		}

		// process the message
		fmt.Printf("Received message: %s\n", string(msg.Value))
		// ... take action based on the message
	}
}
