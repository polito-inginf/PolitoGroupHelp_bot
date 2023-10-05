package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

func ConnectAsReader(topic string) (r *kafka.Reader) {

	kafkaBrokerUrls := getBrokerUrls()

	config := kafka.ReaderConfig{
		Brokers:         kafkaBrokerUrls,
		Topic:           topic,
		MinBytes:        10e1,            // 10B
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}
	r = kafka.NewReader(config)
	return r
}
