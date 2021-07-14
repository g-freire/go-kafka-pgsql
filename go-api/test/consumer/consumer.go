package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

const (
	bootstrapServers = "23.97.166.140:9092"
	//ccloudAPIKey     = "admin"
	//ccloudAPISecret  = "admin"
	ccloudAPIKey     = "pt"
	ccloudAPISecret  = "Qdk<u@jH@=J84^n3vrAj9{6J3T2zya&M"
	//ccloudAPIKey     = "its"
	//ccloudAPISecret  = "*EvLqZ{{U7u!CrGbEVUvTGN>92XM**Q`"
)


func main() {

	topic := "INGESTER2"

	// Now consumes the record and print its value...
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":       bootstrapServers,
		"sasl.mechanisms":         "PLAIN",
		//"security.protocol":       "SASL_SSL",
		"security.protocol":       "SASL_PLAINTEXT",
		"sasl.username":           ccloudAPIKey,
		"sasl.password":           ccloudAPISecret,
		"session.timeout.ms":      6000,
		"group.id":                "PT2",
		"auto.offset.reset":       "earliest"})

	if err != nil {
		panic(fmt.Sprintf("Failed to create consumer: %s", err))
	}

	topics := []string{topic}
	consumer.SubscribeTopics(topics, nil)

	for {
		message, err := consumer.ReadMessage(100 * time.Millisecond)
		if err == nil {
			fmt.Printf("consumed from topic %s [%d] at offset %v: "+
				string(message.Value), *message.TopicPartition.Topic,
				message.TopicPartition.Partition, message.TopicPartition.Offset)
		}
	}

	consumer.Close()
}
