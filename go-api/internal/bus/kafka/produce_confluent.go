package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/satori/go.uuid"
	"strconv"
)


func ProduceConfluent(brokers []string, topic string,  key string, value []byte, producer sarama.SyncProducer) {
	u1 := uuid.Must(uuid.NewV4(), nil)

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		//https://docs.confluent.io/platform/current/installation/configuration/producer-configs.html
		"bootstrap.servers":  brokers[0],
		"enable.idempotence": "true",
		"acks": "all",
		"transactional.id": u1,
		"compression.type":"snappy",
		"retries":5,
	})

	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Delivery report handler for produced messages
	go logEventMsgs(p)

	// TRANSACTION ZONE
	ctx := context.Background()
	err = p.InitTransactions(ctx)
	if err != nil {
		panic(err)
	}
		err = p.BeginTransaction()
		if err != nil {
			fmt.Println(err)
		}
		partitionInt, err := strconv.Atoi(key)
		if err != nil {
			fmt.Println(err)
		}
		//random_partition := kafka.PartitionAny
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(partitionInt)},
			Value:          []byte(value),
			Key:            []byte(key),
		}, nil)
		if err != nil {
			fmt.Println(err)
			_ = p.AbortTransaction(ctx)
		}
		err = p.CommitTransaction(ctx)
		if err != nil {
			fmt.Println(err)
			_ = p.AbortTransaction(ctx)
	}
}


func logEventMsgs(p *kafka.Producer){
	for e := range p.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("\n !!! Delivery failed: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %s[%v]@[%d] %s \n",
					*ev.TopicPartition.Topic,
					ev.TopicPartition.Partition,
					ev.TopicPartition.Offset,
					ev.Value)
			}
		}
	}
}