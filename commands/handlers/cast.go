package handlers

import (
	"github.com/DenrianWeiss/taroly/service/bot"
	"github.com/DenrianWeiss/taroly/service/foundry/cast"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Byte4Cmd struct{}

func (Byte4Cmd) Command() string {
	return "4byte"
}

func (Byte4Cmd) HandleCommand(message tgbotapi.Message) {
	// Call cast cmd
	payload := message.CommandArguments()
	if payload == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a payload.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	resp := cast.SigCall(payload)
	reply := tgbotapi.NewMessage(message.Chat.ID, resp)
	reply.ReplyToMessageID = message.MessageID
	_, _ = bot.GetBot().Send(reply)
}

type DecodeCmd struct{}

func (DecodeCmd) Command() string {
	return "4decode"
}

func (DecodeCmd) HandleCommand(message tgbotapi.Message) {
	// Call cast cmd
	payload := message.CommandArguments()
	if payload == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a payload.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	resp := cast.DecodeCall(payload)
	reply := tgbotapi.NewMessage(message.Chat.ID, resp)
	reply.ReplyToMessageID = message.MessageID
	_, _ = bot.GetBot().Send(reply)
}

type EncodeCmd struct{}

func (EncodeCmd) Command() string {
	return "4encode"
}

func (EncodeCmd) HandleCommand(message tgbotapi.Message) {
	// Call cast cmd
	payload := message.CommandArguments()
	if payload == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a payload.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	args := strings.Split(payload, " ")
	if len(args) < 2 {
		resp := cast.EncodeCall(payload)
		reply := tgbotapi.NewMessage(message.Chat.ID, resp)
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	} else {
		fArgs := args[1:]
		resp := cast.EncodeCall(args[0], fArgs...)
		reply := tgbotapi.NewMessage(message.Chat.ID, resp)
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
}
