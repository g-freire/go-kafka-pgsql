package main

import (
	k "event-driven/internal/bus/kafka"
)

var dev = k.Env{
	Brokers: []string{"127.0.0.1:9092", "127.0.0.1:9094"},
	Topic: "qa",
}

const defaultPostgresURI = "postgres://admin:admin@localhost:6543/admin?sslmode=disable"

//var (
//	partition         = kingpin.Flag("partition", "Partition number").Default("0").String()
//	offsetType        = kingpin.Flag("offsetType", "Offset Type (OffsetNewest | OffsetOldest)").Default("-1").Int()
//	messageCountStart = kingpin.Flag("messageCountStart", "Message counter start from:").Int()
//)


func main() {
	// LOCAL KAFKA
	//k.CreateTopic( dev.Topic, 1, 1)
	//k.SendKeepAliveSignal(dev.Brokers, dev.Topic)
	k.SendKeepAliveSignalLoop(dev.Brokers, dev.Topic, 0)
	//k.StartConsumer(dev.Brokers, dev.Topic, "0",-1, 0)


	// LOCAL POSTGRES DB
	//e.StartLoadTest(defaultPostgresURI)
	//e.StartLoadTestPool(defaultPostgresURI)
	//e.StartLoadTestSessions(defaultPostgresURI)
}

