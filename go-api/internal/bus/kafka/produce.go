package kafka

import (
	"github.com/Shopify/sarama"
	"log"
)


func Produce(topic string,  key string,value []byte, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		//Partition: key,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(value),
		//Value:     sarama.StringEncoder(value),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("\n Message is stored in topic (%s)/ partition(%d) / offset(%d) \n", topic, partition, offset)
}