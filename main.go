package main

import (
	"github.com/DenrianWeiss/taroly/commands/handlers"
	"github.com/DenrianWeiss/taroly/service/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	for update := range bot.GetBot().GetUpdatesChan(tgbotapi.UpdateConfig{}) {

		if update.Message == nil {
			continue
		}
		go handlers.DispatchCommand(*update.Message)
	}
}
