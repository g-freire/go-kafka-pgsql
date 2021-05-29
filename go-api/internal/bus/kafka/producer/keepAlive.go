package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func SendKeepAliveSignal(delay int, topic string, producer sarama.SyncProducer) {
	jobs := make(chan string)
	done := make(chan bool)

	go func() {

		for {
			j, more := <-jobs
			if more {
				msg := time.Now().String()
				produce(topic, "KeepAlive", []byte(msg), producer)
				fmt.Println("KeepAliveMessage Timestamp", j)

			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}

		}

	}()

	for j := 1; j >= 1; j++ {
		t := time.Now().String()
		jobs <- t
		// close signal logic here
		//if j == 3 {
		//	break
		//}

		time.Sleep(time.Duration(delay) * time.Second)
	}
	close(jobs)
	<-done
	fmt.Println("channel is closed")
}
