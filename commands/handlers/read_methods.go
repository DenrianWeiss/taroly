package handlers

import (
	"fmt"
	"github.com/DenrianWeiss/taroly/service/auth"
	"github.com/DenrianWeiss/taroly/service/bot"
	"github.com/DenrianWeiss/taroly/service/cache/user"
	"github.com/DenrianWeiss/taroly/service/eth_rpc"
	"github.com/DenrianWeiss/taroly/service/foundry/cast"
	"github.com/DenrianWeiss/taroly/utils/hx"
	"github.com/ethereum/go-ethereum/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

type GetBalanceCmd struct{}

func (GetBalanceCmd) Command() string {
	return "getbalance"
}

func (GetBalanceCmd) HandleCommand(message tgbotapi.Message) {
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	args := message.CommandArguments()
	var address string
	if args == "" {
		if user.GetUserAccount(strconv.Itoa(int(message.From.ID))) == "" {
			reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide an account address.")
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		} else {
			address = user.GetUserAccount(strconv.Itoa(int(message.From.ID)))
		}
	} else {
		if !common.IsHexAddress(args) {
			reply := tgbotapi.NewMessage(message.Chat.ID, "Invalid address.")
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		}
	}

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

	// Call getbalance cmd
	b, err := eth_rpc.BalanceOf(rpc, address)
	if err != nil {
		reply := tgbotapi.NewMessage(message.Chat.ID, err.Error())
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	reply := tgbotapi.NewMessage(message.Chat.ID, b)
	reply.ReplyToMessageID = message.MessageID
	_, _ = bot.GetBot().Send(reply)
}

type CallCmd struct{}

func (CallCmd) Command() string {
	return "call"
}

func (CallCmd) HandleCommand(message tgbotapi.Message) {
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
	result, err := eth_rpc.Call(rpc,
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

type CallDataCmd struct{}

func (CallDataCmd) Command() string {
	return "calldata"
}

func (CallDataCmd) HandleCommand(message tgbotapi.Message) {
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
	if len(argList) < 2 {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a contract address and payload.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	// Do rpc call.
	result, err := eth_rpc.Call(rpc,
		user.GetUserAccount(strconv.Itoa(int(message.From.ID))),
		argList[0], argList[1], "0")
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

