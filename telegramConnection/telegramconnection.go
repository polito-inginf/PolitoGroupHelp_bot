package telegramconnection

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	godotenv "github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	messagedecoder "PolitoGroupHelpBot/messagedecoder"
)

func Main() {
	err := godotenv.Load(".env")

	// bot setup
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	// message decoder connection
	decoderConn, err := grpc.Dial(":9111", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer decoderConn.Close()

	decoder := messagedecoder.NewMessageDecoderClient(decoderConn)

	for update := range updates {
		if update.Message == nil {
			// see here for all types of messages:
			// https://github.com/go-telegram-bot-api/telegram-bot-api/blob/master/types.go
			continue
		}

		jsonWithTab, _ := json.MarshalIndent(update, "", "\t")
		fmt.Println(string(jsonWithTab))

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
