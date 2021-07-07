package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func StartConfluentConsumerGroup(brokers []string, groupID string, partition string, offsetType, messageCountStart int, consumerID int) {
	log.Println("Starting a new Confluent Consumer Group")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9096",
		"group.id":          "ConfluentTest",
		"auto.offset.reset": "earliest",
		"partition.assignment.strategy": "cooperative-sticky",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"POC2", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()

}
