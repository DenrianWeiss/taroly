package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"sync"
)

var bot *tgbotapi.BotAPI
var botOnce sync.Once

func createBot() {
	token, e := os.LookupEnv("TAROLY_TELEGRAM_TOKEN")
	if !e {
		panic("TAROLY_TELEGRAM_TOKEN not set")
	}
	botI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	bot = botI
}

func GetBot() *tgbotapi.BotAPI {
	botOnce.Do(createBot)
	return bot
}
