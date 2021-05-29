package producer

import (
	"github.com/Shopify/sarama"
	"log"
)

func produce(topic string, key string, value []byte, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: 1,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.StringEncoder(value),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("\n Message is stored in topic (%s)/ partition(%d) / offset(%d) \n", topic, partition, offset)
}

func PublishToKafka(brokers []string, topic string) {

	//var delayBetweenMsg int = 5
	//var deklayBetweenTests int = 20

	//byteValuesFromJSONTest := ReadJsonFiles()
	config := CreatePublisherConfig()

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	//go SendKeepAliveSignal(10, topic, producer)
	SendKeepAliveSignal(1, topic, producer)


	//for {
	//	for i, v := range byteValuesFromJSONTest {
	//		fmt.Printf("\n----------------------------------------------")
	//		fmt.Printf("\n INDEX %d JSON FILE:\n %s \n", i, string(v))
	//		fmt.Printf("----------------------------------------------")
	//
	//		fmt.Printf("\n SLEEPING FOR %d SECONDS\n", delayBetweenMsg)
	//		time.Sleep(time.Duration(delayBetweenMsg) * time.Second)
	//		produce(topic, strconv.Itoa(i), v, producer)
	//	}
	//	fmt.Printf("\n SLEEPING FOR %d SECONDS BEFORE NEXT TEST\n", delayBetweenTests)
	//	time.Sleep(time.Duration(delayBetweenTests) * time.Second)
	//}

}
