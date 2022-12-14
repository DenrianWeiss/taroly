package handlers

import (
	"github.com/DenrianWeiss/taroly/service/auth"
	"github.com/DenrianWeiss/taroly/service/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type AuthCmd struct{}

func (AuthCmd) Command() string {
	return "auth"
}

func (AuthCmd) HandleCommand(message tgbotapi.Message) {
	// Get uid.
	uid := message.From.ID
	// Check power level.
	if !auth.IsRoot(strconv.Itoa(int(uid))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	} else {
		var authCode string
		replyTo := message.ReplyToMessage
		if replyTo != nil {
			authCode = strconv.Itoa(int(replyTo.From.ID))
		} else {
			// Get auth code.
			authCode = message.CommandArguments()
		}
		// Check auth code.
		if authCode == "" {
			reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide an auth id.")
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
		} else {
			auth.AddAuth(authCode)
			reply := tgbotapi.NewMessage(message.Chat.ID, "Auth code added for "+authCode+".")
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
		}
	}
}

type IdCmd struct{}

func (IdCmd) Command() string {
	return "id"
}

func (IdCmd) HandleCommand(message tgbotapi.Message) {
	reply := tgbotapi.NewMessage(message.Chat.ID, strconv.Itoa(int(message.From.ID)))
	reply.ReplyToMessageID = message.MessageID
	_, _ = bot.GetBot().Send(reply)
}
