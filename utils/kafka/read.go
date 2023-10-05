package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Read(reader *kafka.Reader) (res []byte, err error) {
	message, err := reader.ReadMessage(context.Background())
	if err != nil {
		return res, fmt.Errorf("error while receiving message: %s", err.Error())
	}

	value := message.Value
	return value, nil
}
