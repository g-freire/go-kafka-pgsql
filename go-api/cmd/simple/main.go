package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strconv"
)

const (
	bootstrapServers = "13.94.185.238"
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


func main() {

	topic := "confluent-2"
	//createTopic(topic)

	// Produce a new record to the topic...
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to create producer: %s", err))
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
			Partition: kafka.PartitionAny,
		},
		Key: []byte(topic),
		Value: []byte(strconv.Itoa(0))}, nil)
	producer.Close()

	logEventMsgs(producer)

	//var i = 0
	//for {
	//	value := i
	//
	//	producer.Produce(&kafka.Message{
	//		TopicPartition: kafka.TopicPartition{
	//			Topic: &topic,
	//			Partition: kafka.PartitionAny,
	//		},
	//		Key: []byte(topic),
	//		Value: []byte(strconv.Itoa(value))}, nil)
	//
	//	time.Sleep(time.Duration(60*i) * time.Second)
	//	//time.Sleep(100 * time.Millisecond)
	//	fmt.Println("--------------------------------")
	//	fmt.Println("\n ITERATION",i)
	//	fmt.Println("--------------------------------")
	//	i++
	//
	//}
		producer.Close()
}
