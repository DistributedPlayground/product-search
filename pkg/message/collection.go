package message

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Collection interface {
	Listen() error
}

type collection struct {
	kc    *kafka.Consumer
	topic string
}

func NewCollection(kc *kafka.Consumer) Collection {
	return &collection{kc: kc, topic: "collection"}
}
func (c collection) Listen() error {
	// subscribe to a topic
	err := c.kc.Subscribe(c.topic, nil)
	if err != nil {
		panic("Failed to subscribe to topic: " + err.Error())
	}

	// continuously poll for new messages
	for {
		msg, err := c.kc.ReadMessage(-1)
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
