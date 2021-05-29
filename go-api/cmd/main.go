package main

import (
	k "event-driven/internal/bus/kafka"
	"event-driven/internal/bus/kafka/producer"
)

var dev = k.Env{
	Brokers: []string{"127.0.0.1:9092"},
	Topic: "loop_test",
}

func main() {
	// LOCAL KAFKA
	//producer.CreateTopic(dev.Brokers, dev.topic, 1, 2)
	//producer.PublishToKafka(dev.Brokers, dev.Topic)
	//producer.SendKeepAliveSignal(dev.Brokers, dev.Topic)
	producer.SendKeepAliveSignalLoop(dev.Brokers, dev.Topic, 0)


	// LOCAL POSTGRES DB
	//e.StartLoadTest()
	//e.StartLoadTestSessions()

}

