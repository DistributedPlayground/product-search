package message

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DistributedPlayground/product-search/graph/model"
	"github.com/DistributedPlayground/product-search/pkg/repository"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.mongodb.org/mongo-driver/bson"
)

type Collection interface {
	Listen() error
}

type collection struct {
	kc    *kafka.Consumer
	repo  repository.Collection
	topic string
}

func NewCollection(kc *kafka.Consumer, repo repository.Collection) Collection {
	return &collection{kc: kc, repo: repo, topic: "collection"}
}
func (c collection) Listen() error {
	ctx := context.Background()

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

		collection := model.Collection{}
		err = json.Unmarshal(msg.Value, &collection)
		if err != nil {
			fmt.Printf("Error unmarshalling message: %s\n", err)
			continue
		}

		// process the message
		fmt.Printf("Received message: %s\n", string(msg.Value))
		for _, header := range msg.Headers {
			if string(header.Key) == "MessageType" {
				switch string(header.Value) {
				case "Create":
					c.repo.InsertOne(ctx, &collection)
				case "Update":
					c.repo.UpdateOne(ctx, bson.M{"id": collection.ID}, bson.M{"$set": collection})
				case "Delete":
					fmt.Println("Delete")
				default:
					fmt.Println("Unknown")
				}
			}
		}
	}
}
