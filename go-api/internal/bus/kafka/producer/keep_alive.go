package producer

import (
	"log"
	"strconv"
	"time"
)

func SendKeepAliveSignal(brokers []string, topic string) {
	producer := CreateSyncProducer(brokers)
	msg := time.Now().String()
	Produce(topic, "KeepAlive", []byte(msg), producer)
	log.Printf("KeepAliveMessage Timestamp", msg)
	producer.Close()
}

func SendKeepAliveSignalLoop(brokers []string, topic string, delay int) {
	producer := CreateSyncProducer(brokers)
	msg := time.Now().String()
	var i = 0
	for{
		Produce(topic, "KeepAlive" + strconv.Itoa(i), []byte(msg), producer)
		log.Printf("KeepAliveMessage iteration: ", i)
		i++
		time.Sleep(time.Duration(delay) * time.Second)
	}
	producer.Close()
}
