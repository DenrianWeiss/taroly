package tasks

import (
	"fmt"
	"github.com/DenrianWeiss/taroly/service/bot"
	"github.com/DenrianWeiss/taroly/service/db"
	"github.com/DenrianWeiss/taroly/service/foundry/cast"
	"github.com/DenrianWeiss/taroly/utils/hx"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

func TraceJob(chatId int64, messageId int, rpc, txId string) {
	// Call cast
	call := cast.RunCall(rpc, txId)
	doc := hx.HandleTerminalEscape(call)
	// Save call to db
	db.Set(db.GetDb(), []byte("trace"+txId), []byte(doc))
	// Send result to user
	baseUri, _ := os.LookupEnv("TAROLY_WEB_URL")
	reply := tgbotapi.NewMessage(chatId, fmt.Sprintf("Trace result for %s:\n%strace/%s", txId, baseUri, txId))
	reply.ReplyToMessageID = messageId
	_, _ = bot.GetBot().Send(reply)
}
