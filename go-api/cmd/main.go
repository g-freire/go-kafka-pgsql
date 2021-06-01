package main

import (
	k "event-driven/internal/bus/kafka"
	l "event-driven/pkg/load_tests"
)

var dev = k.Env{
	Brokers: []string{"127.0.0.1:9092"},
	Topic:   "bet",
}

const defaultPostgresURI = "postgres://admin:admin@localhost:6543/admin?sslmode=disable"

func main() {
	// LOCAL KAFKA
	//k.CreateTopic( dev.Topic, 1, 2)
	//k.SendKeepAliveSignal(dev.Brokers, dev.Topic)
	//producer.SendKeepAliveSignalLoop(dev.Brokers, dev.Topic, 0)
	//k.StartConsumer(dev.Brokers, dev.Topic, "0",-1, 0)

	// LOCAL POSTGRES DB
	//l.StartLoadTest(defaultPostgresURI)
	l.StartLoadTestPool(defaultPostgresURI)
	//l.StartLoadTestSessions(defaultPostgresURI)
}
