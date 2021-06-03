package main

import (
	k "event-driven/internal/bus/kafka"
)

var dev = k.Env{
	Brokers: []string{"127.0.0.1:9092", "127.0.0.1:9094"},
	Topic: "qa",
}

const defaultPostgresURI = "postgres://admin:admin@localhost:6543/admin?sslmode=disable"

func main() {
	// LOCAL KAFKA
	//k.CreateTopic( dev.Topic, 1, 1)
	//k.SendKeepAliveSignal(dev.Brokers, dev.Topic)
	//k.SendKeepAliveSignalLoop(dev.Brokers, dev.Topic, 0)
	k.StartConsumer(dev.Brokers, dev.Topic, "0",-1, 0)

	// LOCAL POSTGRES DB
	//e.StartLoadTest(defaultPostgresURI)
	//e.StartLoadTestPool(defaultPostgresURI)
	//e.StartLoadTestSessions(defaultPostgresURI)
	//l.StartLoadTest(defaultPostgresURI)
	//l.StartLoadTestPool(defaultPostgresURI)
	//l.StartLoadTestSessions(defaultPostgresURI)
}
