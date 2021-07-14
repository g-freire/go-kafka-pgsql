package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	bootstrapServers = "23.97.166.140:9092"
	ccloudAPIKey     = "admin"
	ccloudAPISecret  = "admin"
	//ccloudAPIKey     = "pt"
	//ccloudAPISecret  = "Qdk<u@jH@=J84^n3vrAj9{6J3T2zya&M"
	//ccloudAPIKey     = "its"
	//ccloudAPISecret  = "*EvLqZ{{U7u!CrGbEVUvTGN>92XM**Q`"


)

func logEventMsgs(p *kafka.Producer){
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("\n !!! Delivery failed: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %s[%v]@[%d] %s \n",
					*ev.TopicPartition.Topic,
					ev.TopicPartition.Partition,
					ev.TopicPartition.Offset,
					ev.Value)
			}
		}
	}
}

func signalLoop(value string, topic string, producer *kafka.Producer) {
	for {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic,
				Partition: kafka.PartitionAny},
			Value: []byte(value)}, nil)
		//time.Sleep(.1 * time.Second)
	}
}

func receiveLog(producer *kafka.Producer){
	// Wait for delivery report
	e := <-producer.Events()

	message := e.(*kafka.Message)
	if message.TopicPartition.Error != nil {
		fmt.Printf("failed to deliver message: %v\n",
			message.TopicPartition)
	} else {
		fmt.Printf("delivered to topic %s [%d] at offset %v\n",
			*message.TopicPartition.Topic,
			message.TopicPartition.Partition,
			message.TopicPartition.Offset)
	}
	//producer.Close()
}

func main() {

	topic := "INGESTER2"
	//createTopic(topic)

	// Produce a new record to the topic...
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"sasl.mechanisms":   "PLAIN",
		//"security.protocol":       "SASL_SSL",
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.username":     ccloudAPIKey,
		"sasl.password":     ccloudAPISecret})

	if err != nil {
		panic(fmt.Sprintf("Failed to create producer: %s", err))
	}

	go logEventMsgs(producer)
	signalLoop("golang test value", topic, producer)


}
