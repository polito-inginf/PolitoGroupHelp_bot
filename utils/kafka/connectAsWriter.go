package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

func ConnectAsWriter(topic string) (w *kafka.Writer, err error) {
	// func ConnectAsWriter(clientId string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		// ClientID: clientId,
	}

	kafkaBrokerUrls := getBrokerUrls()
	
	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
	}
	w = kafka.NewWriter(config)
	return w, nil
}
