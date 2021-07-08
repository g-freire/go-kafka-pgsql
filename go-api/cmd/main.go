package main

import 		(
	kafka "event-driven/internal/bus/kafka"
	"flag"
	"log"
)

var dev = kafka.Env{
	Brokers: []string{"127.0.0.1:9092", "127.0.0.1:9094"},
	//Brokers: []string{"127.0.0.1:9092", "127.0.0.1:9094", "127.0.0.1:9096"},
	Topic: "INGESTER",
	//Topic: "POC2",
}

const defaultPostgresURI = "postgres://admin:admin@localhost:6543/admin?sslmode=disable"

func main() {
	// CLI
	kafka_type := flag.String("t", "c", "Types of kafka eg. -t=c for consumer or -t=p for producer")
	id := flag.Int("id", 1, "ID of the consumer or producer ")
	flag.Parse()
	log.Println("  KAFKA TYPE:", *kafka_type, "ID:", *id)

	//////////////////////////////////////////////////////////////////////////////////////////////
	// KAFKA
	//////////////////////////////////////////////////////////////////////////////////////////////
	//kafka.CreateTopic(dev.Topic, 4, 1, dev.Brokers[1])
	//return

	if *kafka_type == "p" {
		log.Println("  KAFKA TYPE PRODUCER :", *kafka_type)
		//k.SendKeepAliveSignal(dev.Brokers, dev.Topic)
		kafka.SendKeepAliveSignalLoop(dev.Brokers, dev.Topic, 1, *id)
	}
	if *kafka_type == "c" {
		log.Println("  KAFKA TYPE CONSUMER :", *kafka_type)
		//k.StartConsumer(dev.Brokers, dev.Topic, "0",-1, 0, *id)
		// kafka.StartConsumerGroup(dev.Brokers, dev.Topic, "0",-1, 0, *id)
		kafka.StartConfluentConsumerGroup(dev.Brokers, []string{dev.Topic}, "TWorkers", string(2), 0, 0, *id)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////
	// POSTGRES DB LOAD TESTS
	//////////////////////////////////////////////////////////////////////////////////////////////
	//e.StartLoadTest(defaultPostgresURI)
	// e.StartLoadTestPool(defaultPostgresURI)
	//e.StartLoadTestSessions(defaultPostgresURI)
	//l.StartLoadTest(defaultPostgresURI)
	//l.StartLoadTestPool(defaultPostgresURI)
	//l.StartLoadTestSessions(defaultPostgresURI)
}
