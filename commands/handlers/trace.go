package handlers

import (
	"github.com/DenrianWeiss/taroly/commands/tasks"
	"github.com/DenrianWeiss/taroly/service/auth"
	"github.com/DenrianWeiss/taroly/service/bot"
	"github.com/DenrianWeiss/taroly/service/cache/user"
	"github.com/DenrianWeiss/taroly/utils/hx"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

type TraceCmd struct{}

func (TraceCmd) Command() string {
	return "trace"
}

func (TraceCmd) HandleCommand(message tgbotapi.Message) {
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	arg := message.CommandArguments()
	argSplit := strings.Split(arg, " ")
	if argSplit[0] == "" || !hx.IsValidHex(argSplit[0]) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a transaction hash.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	// Get RPC
	rpc := user.GetUserRpcUrl(strconv.Itoa(int(message.From.ID)))
	if rpc == "" {
		if user.GetUserOnlineMode(strconv.Itoa(int(message.From.ID))) {
			reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a chain name for online mode.")
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		} else {
			reply := tgbotapi.NewMessage(message.Chat.ID, "Please active a fork for offline mode.")
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		}
	}
	if user.GetUserOnlineMode(strconv.Itoa(int(message.From.ID))) && len(argSplit) == 1 {
		reply := tgbotapi.NewMessage(message.Chat.ID, "For online mode please use tenderly. https://dashboard.tenderly.co/explorer")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	} else {
		// Offline mode
		reply := tgbotapi.NewMessage(message.Chat.ID, "this will take somme minutes.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		go tasks.TraceJob(message.Chat.ID, message.MessageID, rpc, argSplit[0])
	}
}
