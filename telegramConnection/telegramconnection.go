package telegramconnection

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	messagedecoder "PolitoGroupHelpBot/messagedecoder"
	"PolitoGroupHelpBot/utils"
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

	// message decoder connection
	decoderConn, err := grpc.Dial(utils.LoadPortFromEnv("DECODER_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer decoderConn.Close()

	// handle messages
	decoder := messagedecoder.NewMessageDecoderClient(decoderConn)

	for update := range updates {
		if update.Message == nil {
			// see here for all types of messages:
			// https://github.com/go-telegram-bot-api/telegram-bot-api/blob/master/types.go
			continue
		}

		jsonNoTab, _ := json.MarshalIndent(update, "", "")
		res, err := decoder.Decode(context.Background(), &messagedecoder.TgMessageInfo{MessageInfo: string(jsonNoTab)})
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		fmt.Println(res)
	}
}

func NewMessageDecoderClient(decoderConn *grpc.ClientConn) {
	panic("unimplemented")
}
