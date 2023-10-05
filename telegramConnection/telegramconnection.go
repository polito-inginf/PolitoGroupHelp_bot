package telegramconnection

import (
	"context"
	"encoding/json"
	"log"

	"PolitoGroupHelpBot/utils"
	"PolitoGroupHelpBot/utils/kafka"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
)

func Main() {
	// bot setup
	bot, err := tgbotapi.NewBotAPI(utils.LoadEnv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Fatalf("Could not connect to Telegram bot: %v", err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	queueWriter, err := kafka.ConnectAsWriter("before-filtering")
	if err != nil {
		log.Fatalf("Could not connect to Kafka: %v", err)
	}
	defer queueWriter.Close()

	log.Default().Println("Connected to the Telegram API, listening")

	for update := range updates {
		if update.Message == nil {
			// see here for all types of messages:
			// https://github.com/go-telegram-bot-api/telegram-bot-api/blob/master/types.go
			continue
		}

		jsonNoTab, _ := json.MarshalIndent(update, "", "")
		kafka.Write(context.Background(), []byte("test"), []byte(jsonNoTab), queueWriter)

		log.Default().Println("Message received and stored to the queue")
	}
}

func NewMessageDecoderClient(decoderConn *grpc.ClientConn) {
	panic("unimplemented")
}
