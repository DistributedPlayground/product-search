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

type Product interface {
	Listen() error
}

type product struct {
	kc    *kafka.Consumer
	repo  repository.Product
	topic string
}

func NewProduct(kc *kafka.Consumer, repo repository.Product) Product {
	return &product{kc: kc, repo: repo, topic: "product"}
}
func (p product) Listen() error {
	ctx := context.Background()

	// subscribe to a topic
	err := p.kc.Subscribe(p.topic, nil)
	if err != nil {
		panic("Failed to subscribe to topic: " + err.Error())
	}

	// continuously poll for new messages
	for {
		msg, err := p.kc.ReadMessage(-1)
		if err != nil {
			// handle error
			fmt.Printf("Error reading message: %s\n", err)
			continue
		}

		product := model.Product{}
		err = json.Unmarshal(msg.Value, &product)
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
					p.repo.InsertOne(ctx, &product)
				case "Update":
					p.repo.UpdateOne(ctx, bson.M{"id": product.ID}, bson.M{"$set": product})
				case "Delete":
					fmt.Println("Delete")
				default:
					fmt.Println("Unknown")
				}
			}
		}
	}
}
