package main

import (
	k "event-driven/internal/bus/kafka"
	e "event-driven/pkg/event"
)

var dev = k.Env{
	Brokers: []string{"127.0.0.1:9092"},
	Topic: "loop_test",
}

const defaultPostgresURI = "postgres://admin:admin@localhost:6543/admin?sslmode=disable"


func main() {
	// LOCAL KAFKA
	//producer.CreateTopic(dev.Brokers, dev.topic, 1, 2)
	//producer.SendKeepAliveSignal(dev.Brokers, dev.Topic)
	//producer.SendKeepAliveSignalLoop(dev.Brokers, dev.Topic, 0)


	// LOCAL POSTGRES DB
	e.StartLoadTest(defaultPostgresURI)
	//e.StartLoadTestSessions(defaultPostgresURI)

}

