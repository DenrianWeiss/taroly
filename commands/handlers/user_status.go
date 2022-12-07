package handlers

import (
	"fmt"
	"github.com/DenrianWeiss/taroly/service/auth"
	"github.com/DenrianWeiss/taroly/service/bot"
	"github.com/DenrianWeiss/taroly/service/cache/rpc"
	"github.com/DenrianWeiss/taroly/service/cache/user"
	"github.com/ethereum/go-ethereum/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type OnlineModeCmd struct{}

func (OnlineModeCmd) Command() string {
	return "onlinemode"
}

func (OnlineModeCmd) HandleCommand(message tgbotapi.Message) {
	// Check user power level.
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	args := message.CommandArguments()
	if args == "true" {
		user.SetUserOnlineMode(strconv.Itoa(int(message.From.ID)), true)
		reply := tgbotapi.NewMessage(message.Chat.ID, "Online mode enabled.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	} else if args == "false" {
		user.SetUserOnlineMode(strconv.Itoa(int(message.From.ID)), false)
		reply := tgbotapi.NewMessage(message.Chat.ID, "Online mode disabled.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	} else {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a boolean argument.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	}
}

type SetChainCmd struct{}

func (SetChainCmd) Command() string {
	return "setchain"
}

func (SetChainCmd) HandleCommand(message tgbotapi.Message) {
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	args := message.CommandArguments()
	if rpc.GetRpcUrl(args) == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a valid chain name.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	} else {
		user.SetUserChain(strconv.Itoa(int(message.From.ID)), args)
		reply := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Chain set to %s.", args))
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	}
}

type GetChainCmd struct{}

func (GetChainCmd) Command() string {
	return "getchain"
}

func (GetChainCmd) HandleCommand(message tgbotapi.Message) {
	chains := rpc.ListChains()
	reply := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Available chains: %s", chains))
	reply.ReplyToMessageID = message.MessageID
	_, _ = bot.GetBot().Send(reply)
}

type SetAccountCmd struct{}

func (SetAccountCmd) Command() string {
	return "setaccount"
}

func (SetAccountCmd) HandleCommand(message tgbotapi.Message) {
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	args := message.CommandArguments()
	v := common.IsHexAddress(args)
	if !v {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a valid account address.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	}
	user.SetUserAccount(strconv.Itoa(int(message.From.ID)), args)
	reply := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Account set to %s.", args))
	reply.ReplyToMessageID = message.MessageID
	_, _ = bot.GetBot().Send(reply)
}

type GetAccountCmd struct{}

func (GetAccountCmd) Command() string {
	return "getaccount"
}

func (GetAccountCmd) HandleCommand(message tgbotapi.Message) {
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	userAccount := user.GetUserAccount(strconv.Itoa(int(message.From.ID)))
	if userAccount == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please set your account first.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	} else {
		reply := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Your account is %s.", userAccount))
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	}
}
