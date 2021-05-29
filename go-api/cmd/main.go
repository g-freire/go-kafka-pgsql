package main

import (
	k "event-driven/internal/bus/kafka"
	e "event-driven/pkg/event"
)

var dev = k.Env{
	Brokers: []string{"127.0.0.1:9092"},
	Topic: "loop_test",
}

func main() {
	// LOCAL KAFKA
	//producer.CreateTopic(dev.topic, 1, 2)
	//producer.PublishToKafka(dev.Brokers, dev.Topic)

	// LOCAL POSTGRES DB
	e.StartLoadTest()

}
