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
	Produce(topic, "1", []byte(msg), producer)
	//Produce(topic, "KeepAlive", []byte(msg), producer)
	log.Printf("KeepAliveMessage Timestamp", msg)
	producer.Close()
}

type Response struct{
	ProducerID int `json:"producer_id"`
	ProducerTimestamp string `json:"producer_timestamp"`
	Value int `json:"value"`

}

func SendKeepAliveSignalLoop(brokers []string, topic string, delay int, producerID int) {
	//producer := CreateSyncProducer(brokers)

	var i = 0
	producerIDString := strconv.Itoa(producerID)

	for{
		r := Response{
			producerID,
			time.Now().String(),
			i,
		}
		b, err := json.Marshal(r)
		if err != nil {
			fmt.Println("error marshelling:", err)
		}

		//Produce(topic, "KeepAlive" + strconv.Itoa(i), b, producer)
		ProduceConfluent(brokers,topic, producerIDString, b)

		log.Printf("KeepAliveMessage iteration: ", i)
		fmt.Println("--------------------------------")
		fmt.Println("\n ITERATION",i)
		fmt.Println("--------------------------------")
		i++
		time.Sleep(time.Duration(60 * i) * time.Second)
	}
	//producer.Close()
}
