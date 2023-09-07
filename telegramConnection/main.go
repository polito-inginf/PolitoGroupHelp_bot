package main

import (
	"fmt"
	"os"
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	godotenv "github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))

	if err != nil {
		panic(err)
	}

	// bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			// see here for all types of messages:
			// https://github.com/go-telegram-bot-api/telegram-bot-api/blob/master/types.go
			continue
		}

		json, _ := json.MarshalIndent(update, "", "\t" )
		fmt.Println(string(json))
	}
}
