package kafka

import (
	"PolitoGroupHelpBot/utils"
	"fmt"
)

func getBrokerUrls() (kafkaBrokerUrls []string) {

	kafkaBaseUrl := utils.LoadEnv("KAFKA_URL")

	for i := 1; ; i++ {
		port, err := utils.LoadEnvIfPresent(fmt.Sprintf("KAFKA_%v_PORT", i))
		if err != nil {
			break
		}

		kafkaBrokerUrls = append(kafkaBrokerUrls, fmt.Sprintf("%v:%v", kafkaBaseUrl, port))
	}

	if len(kafkaBrokerUrls) == 0 {
		panic("No Kafka brokers present")
	}

	return
}
