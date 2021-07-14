package kafka

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartConfluentConsumerGroup(brokers []string, topics []string, groupID string, partition string, offsetType, messageCountStart int, consumerID int) {
	log.Println("Starting a new Confluent Consumer Group")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":             brokers[2],
		"group.id":                      groupID,
		//"auto.offset.reset":             "earliest",
		"partition.assignment.strategy": "roundrobin",
		"enable.auto.commit":            "false",
		"isolation.level":               "read_committed",
	})

	if err != nil {
		panic(err)
	}

	// c.SubscribeTopics([]string{"POC2", "POC2"}, nil)
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			// c.Commit()
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()

}
