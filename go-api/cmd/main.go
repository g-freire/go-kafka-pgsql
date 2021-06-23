package main

import (
	k "event-driven/internal/bus/kafka"
	"flag"
	"log"
)

var dev = k.Env{
	Brokers: []string{"127.0.0.1:9092", "127.0.0.1:9094"},
	Topic: "qa",
}

const defaultPostgresURI = "postgres://admin:admin@localhost:6543/admin?sslmode=disable"



func main() {
	// CLI
	kafka_type := flag.String("t", "p", "Types of kafka eg. -t=c for consumer or -t=p for producer")
	id := flag.Int("id", 1, "ID of the consumer or producer ")

	flag.Parse()
	log.Println("\nKAFKA TYPE:", *kafka_type)

	// LOCAL KAFKA
	//k.CreateTopic( dev.Topic, 1, 1)
	if *kafka_type == "p"{
		log.Println("\nKAFKA TYPE PRODUCER :", *kafka_type)
		//k.SendKeepAliveSignal(dev.Brokers, dev.Topic)
		k.SendKeepAliveSignalLoop(dev.Brokers, dev.Topic, 10, *id)
	}
	if *kafka_type == "c"{
		log.Println("\nKAFKA TYPE CONSUMER :", *kafka_type)
		k.StartConsumer(dev.Brokers, dev.Topic, "0",-1, 0, *id)
	}



	// LOCAL POSTGRES DB
	//e.StartLoadTest(defaultPostgresURI)
	// e.StartLoadTestPool(defaultPostgresURI)
	//e.StartLoadTestSessions(defaultPostgresURI)
	//l.StartLoadTest(defaultPostgresURI)
	//l.StartLoadTestPool(defaultPostgresURI)
	//l.StartLoadTestSessions(defaultPostgresURI)
}