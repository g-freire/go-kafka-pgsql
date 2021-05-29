package producer

import (
	"github.com/Shopify/sarama"
	"log"
)

func CreateSyncProducer(brokers []string) sarama.SyncProducer {

	config := CreatePublisherConfigStruct()

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Panic(err)
	}
	//defer func() {
	//	if err := producer.Close(); err != nil {
	//		log.Panic(err)
	//	}
	//}()

	return producer
}
