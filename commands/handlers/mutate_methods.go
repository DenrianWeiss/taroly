package handlers

import (
	"fmt"
	"github.com/DenrianWeiss/taroly/service/auth"
	"github.com/DenrianWeiss/taroly/service/bot"
	"github.com/DenrianWeiss/taroly/service/cache/user"
	"github.com/DenrianWeiss/taroly/service/eth_rpc"
	"github.com/DenrianWeiss/taroly/service/foundry/cast"
	"github.com/DenrianWeiss/taroly/utils/hx"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

type SetBalanceCmd struct{}

func (SetBalanceCmd) Command() string {
	return "set_balance"
}

func (SetBalanceCmd) HandleCommand(message tgbotapi.Message) {
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	if user.GetUserOnlineMode(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "SetBalance is not supported in online mode.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	if user.GetUserRpcUrl(strconv.Itoa(int(message.From.ID))) == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please active a fork for offline mode.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	if user.GetUserAccount(strconv.Itoa(int(message.From.ID))) == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide an account address.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	args := message.CommandArguments()
	if args == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a balance number.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	rpc := user.GetUserRpcUrl(strconv.Itoa(int(message.From.ID)))
	address := user.GetUserAccount(strconv.Itoa(int(message.From.ID)))
	balance, err := eth_rpc.SetBalance(rpc, address, args)
	if err != nil {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Error: "+err.Error())
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	}
	reply := tgbotapi.NewMessage(message.Chat.ID, "Set balance successfully."+balance)
	reply.ReplyToMessageID = message.MessageID
	_, _ = bot.GetBot().Send(reply)
}

type TransactCmd struct{}

func (TransactCmd) Command() string {
	return "transact"
}

func (TransactCmd) HandleCommand(message tgbotapi.Message) {
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	if user.GetUserAccount(strconv.Itoa(int(message.From.ID))) == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide an account address using /setaccount.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	args := message.CommandArguments()
	if args == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a contract address.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	argList := strings.Split(args, " ")
	if len(argList) < 1 {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a contract address and a function name.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	// Get RPC Address
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
	var argsCall []string
	if len(argList) > 2 {
		argsCall = argList[2:]
	}

	var abi string
	// Use Cast to encode the function call
	if len(argList) >= 1 {
		abi = cast.EncodeCall(argList[1], argsCall...)
		// See if the result is valid call data.
		if !hx.IsValidHex(abi) {
			reply := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Something Wrong happened: %s", abi))
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		}
	} else {
		abi = "0x"
	}
	// Do rpc call.
	// Impersonate the account
	_, err := eth_rpc.Impersonate(rpc, user.GetUserAccount(strconv.Itoa(int(message.From.ID))))
	if err != nil {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Error: "+err.Error())
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	result, err := eth_rpc.Send(rpc,
		user.GetUserAccount(strconv.Itoa(int(message.From.ID))),
		argList[0], abi, "0")
	if err != nil {
		reply := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Error: %s", hx.FilterUnPrintable(err.Error())))
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	} else {
		reply := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Result: %s", result))
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
}
