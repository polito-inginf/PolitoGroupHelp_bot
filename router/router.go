package router

import (
	"PolitoGroupHelpBot/utils"
	"PolitoGroupHelpBot/utils/kafka"
	"log"
)

func Main() {
	topic := ""

	if utils.IsEnvPresent("SKIP_FILTERS") {
		topic = "before-filtering"
	} else {
		topic = "after-filtering"
	}

	reader := kafka.ConnectAsReader(topic)
	defer reader.Close()

	log.Default().Printf("Started router, listening on queue %v\n", topic)

	for {
		message, err := kafka.Read(reader)
		if err != nil {
			log.Default().Println(err)
			continue
		}

		log.Default().Println(message)
	}
}
