package kafka

import (
	"log"
	"time"

	"github.com/Shopify/sarama" // Sarama 1.22.0
)

func CreateTopic(topicName string, partitions int32, replication int16) {
	broker := sarama.NewBroker("localhost:9094")
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	broker.Open(config)
	connected, err := broker.Connected()
	if err != nil {
		log.Print(err.Error())
	}
	log.Print("Connected :", connected)

	// Setup the Topic details in CreateTopicRequest struct
	topic := topicName
	topicDetail := &sarama.TopicDetail{}
	topicDetail.NumPartitions = partitions
	topicDetail.ReplicationFactor = replication
	topicDetail.ConfigEntries = make(map[string]*string)

	topicDetails := make(map[string]*sarama.TopicDetail)
	topicDetails[topic] = topicDetail

	request := sarama.CreateTopicsRequest{
		Timeout:      time.Second * 45,
		TopicDetails: topicDetails,
	}

	// Send request to Broker
	response, err := broker.CreateTopics(&request)

	// handle errors if any
	if err != nil {
		log.Printf("%#v", &err)
	}
	t := response.TopicErrors
	for key, val := range t {
		log.Printf("Key is %s", key)
		log.Printf("Value is %#v", val.Err.Error())
		log.Printf("Value3 is %#v", val.ErrMsg)
	}
	log.Printf("the response is %#v", response)

	// close connection to broker
	broker.Close()
}
