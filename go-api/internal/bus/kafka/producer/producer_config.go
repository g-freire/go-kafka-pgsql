package producer

import (
	"github.com/Shopify/sarama"
	"time"
)

func CreatePublisherConfig() *sarama.Config {

	// RETURNS THE SANE DEFAULT CONFIG STRUCT
	config := sarama.NewConfig()

	// AUTH - SALS PLAINTEXT
	// config.Net.SASL.Enable = true
	// config.Net.SASL.User = "admin"
	// config.Net.SASL.Password = "admin"
	// config.Net.SASL.User = "pt"
	// config.Net.SASL.Password = "Qdk<u@jH@=J84^n3vrAj9{6J3T2zya&M"
	// config.Net.SASL.User = "its"
	// config.Net.SASL.Password = "*EvLqZ{{U7u!CrGbEVUvTGN>92XM**Q`"

	// PROD PRODUCER

	// If enabled, the producer will ensure that exactly one copy of each message
	// is written. Requires Net.MaxOpenRequests to be 1
	config.Producer.Idempotent = true
	// The maximum permitted size of a message (defaults to 1000000). Should be
	// set equal to or smaller than the broker's `message.max.bytes`.
	config.Producer.MaxMessageBytes = 2000000
	// (defaults to  no compression)
	config.Producer.Compression = 1 //CompressionGZIP
	// (defaults to WaitForLocal)
	config.Producer.RequiredAcks = sarama.WaitForAll
	// ACKS timeout (defaults to 10 seconds)
	config.Producer.Timeout = 30 * time.Second
	// The total number of times to retry sending a message (default 3).
	config.Producer.Retry.Max = 10
	// How long to wait for leader election to occur before retrying (default 250ms)
	config.Producer.Retry.Backoff = 5 * time.Second
	// (default disabled)
	config.Producer.Return.Successes = true
	// How long to wait for a transmit. ( defaults 30 sec)
	config.Net.WriteTimeout = 40 * time.Second
	config.Net.DialTimeout = 40 * time.Second
	// How many outstanding requests a connection is allowed to have before
	// sending on it blocks (default 5).
	config.Net.MaxOpenRequests = 1
	config.Metadata.Retry.Backoff = 500 * time.Millisecond
	config.Metadata.RefreshFrequency = 8 * time.Minute

	return config

}
