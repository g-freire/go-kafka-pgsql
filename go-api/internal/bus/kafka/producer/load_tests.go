package producer

import (
	"fmt"
	"strconv"
	"time"
)

func RunLoadTestCases(brokers []string, topic string) {

	var delayBetweenMsg int = 5
	var delayBetweenTests int = 20
	byteValuesFromJSONTest := ReadJsonFiles()

	producer := CreateSyncProducer(brokers)
	//go SendKeepAliveSignal(1, topic, producer)

	for {
		for i, v := range byteValuesFromJSONTest {
			fmt.Printf("\n----------------------------------------------")
			fmt.Printf("\n INDEX %d JSON FILE:\n %s \n", i, string(v))
			fmt.Printf("----------------------------------------------")

			fmt.Printf("\n SLEEPING FOR %d SECONDS\n", delayBetweenMsg)
			time.Sleep(time.Duration(delayBetweenMsg) * time.Second)
			Produce(topic, strconv.Itoa(i), v, producer)
		}
		fmt.Printf("\n SLEEPING FOR %d SECONDS BEFORE NEXT TEST\n", delayBetweenTests)
		time.Sleep(time.Duration(delayBetweenTests) * time.Second)
	}
	producer.Close()
}
