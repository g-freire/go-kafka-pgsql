package kafka

import (
	"context"
	"encoding/json"
	db "event-driven/internal/db/postgres"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"time"
)

//var (
//	brokers           = []string{}
//	topic             = kingpin.Flag("topic", "Topic name").Default("qa").String()
//	partition         = kingpin.Flag("partition", "Partition number").Default("0").String()
//	offsetType        = kingpin.Flag("offsetType", "Offset Type (OffsetNewest | OffsetOldest)").Default("-1").Int()
//	messageCountStart = kingpin.Flag("messageCountStart", "Message counter start from:").Int()
//)

const postgresURI = "postgres://admin:admin@localhost:6543/admin?sslmode=disable"


func StartConsumer(brokers []string , topic, partition string, offsetType, messageCountStart int, consumerID int)  {


	//kingpin.Parse()
	config := sarama.NewConfig()

	// SECURITY
	// config.Net.SASL.Enable = true
	// config.Net.SASL.Handshake = true
	// config.Net.SASL.User = "admin"
	// config.Net.SASL.Password = "admin"

	config.Consumer.Return.Errors = true
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			log.Panic(err)
		}
	}()
	//consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Panic(err)
	}

	//INSERT LOGIC
	client := db.NewPostgresConnectionPool(postgresURI)
	defer client.Close()


	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)
			case msg := <-consumer.Messages():
				var m Response
				err := json.Unmarshal(msg.Value, &m)
				if err != nil {
					fmt.Println("error unmarshelling:", err)
				}
				messageCountStart++
				log.Println("Received messages", string(msg.Key), string(msg.Value))
				timeNow := time.Now().String()

				sql := `INSERT INTO KAFKA (producer_id, producer_timestamp,
										   consumer_id, consumer_timestamp) VALUES ($1, $2, $3, $4)`
				_, err = client.Exec(context.Background(), sql, m.ProducerID, m.ProducerTimestamp, consumerID, timeNow)
				if err != nil {
					panic(err)
				} else {
					fmt.Print("\n INSERTED ", client, " ", msg)
				}
			case <-signals:
				log.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	log.Println("Processed", messageCountStart, "messages")
}
