package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func Write(parent context.Context, key, value []byte, writer *kafka.Writer) (err error) {
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	return writer.WriteMessages(parent, message)
}
