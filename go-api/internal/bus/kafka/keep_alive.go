package kafka

import (
	"encoding/json"
	"fmt"
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

type Response struct{
	ProducerID int
	Timestamp string
}

func SendKeepAliveSignalLoop(brokers []string, topic string, delay int, producerID int) {
	producer := CreateSyncProducer(brokers)
	r := Response{
		producerID,
		time.Now().String(),
	}
	b, err := json.Marshal(r)

	//var x, y int
	//m := Multiply{
	//	X: x,
	//	Y: y,
	//}

	//jsonMsg, err := json.Marshal(m)


	if err != nil {
		fmt.Println("error marshelling:", err)
	}
	var i = 0
	for{
		Produce(topic, "KeepAlive" + strconv.Itoa(i), b, producer)
		log.Printf("KeepAliveMessage iteration: ", i)
		i++
		time.Sleep(time.Duration(delay) * time.Second)
	}
	producer.Close()
}
