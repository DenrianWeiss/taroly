package handlers

import (
	json2 "encoding/json"
	"fmt"
	"github.com/DenrianWeiss/taroly/model"
	"github.com/DenrianWeiss/taroly/service/auth"
	"github.com/DenrianWeiss/taroly/service/bot"
	"github.com/DenrianWeiss/taroly/service/cache/rpc"
	"github.com/DenrianWeiss/taroly/service/cache/user"
	"github.com/DenrianWeiss/taroly/service/eth_rpc"
	"github.com/DenrianWeiss/taroly/service/foundry/anvil"
	"github.com/DenrianWeiss/taroly/utils/hmac"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
	"strings"
)

type NewForkCmd struct {
}

func (NewForkCmd) Command() string {
	return "newfork"
}

func (NewForkCmd) HandleCommand(message tgbotapi.Message) {
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	args := message.CommandArguments()
	if args == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a fork name.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	if user.GetUserStatus(strconv.Itoa(int(message.From.ID))).SimulationPid != 0 {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are already running a fork.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	argsList := strings.Split(args, " ")
	if len(argsList) > 2 {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please provide a fork name and a fork block number.\n"+
			"Example: /newfork mainnet 11451419")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	// Test chain validity
	v := rpc.GetRpcUrl(argsList[0])
	if v == "" {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Invalid chain name.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	if len(argsList) == 1 {
		// Start fork
		fork, port, err := anvil.StartFork(argsList[0], 0)
		if err != nil {
			reply := tgbotapi.NewMessage(message.Chat.ID, "Failed to start fork: "+err.Error())
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		} else {
			user.SetUserSimulation(strconv.Itoa(int(message.From.ID)), fork, port)
			reply := tgbotapi.NewMessage(message.Chat.ID, "Fork started.")
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		}
	} else {
		// Validate the block number
		blocknumber, err := strconv.Atoi(argsList[1])
		if err != nil {
			reply := tgbotapi.NewMessage(message.Chat.ID, "Invalid block number.")
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		}
		onlineBlocknumber, err := eth_rpc.BlockNumber(v)
		if int64(blocknumber) > onlineBlocknumber {
			reply := tgbotapi.NewMessage(message.Chat.ID, "Invalid block number.")
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		}
		// Start fork
		fork, port, err := anvil.StartFork(argsList[0], int64(blocknumber))
		if err != nil {
			reply := tgbotapi.NewMessage(message.Chat.ID, "Failed to start fork: "+err.Error())
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		} else {
			user.SetUserSimulation(strconv.Itoa(int(message.From.ID)), fork, port)
			reply := tgbotapi.NewMessage(message.Chat.ID, "Fork started.")
			reply.ReplyToMessageID = message.MessageID
			_, _ = bot.GetBot().Send(reply)
			return
		}
	}
}

type StopForkCmd struct {
}

func (StopForkCmd) Command() string {
	return "stopfork"
}

func (StopForkCmd) HandleCommand(message tgbotapi.Message) {
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}

	s := user.GetUserStatus(strconv.Itoa(int(message.From.ID)))
	if s.SimulationPid == 0 {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not running a fork.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	// Stop Fork
	err := anvil.StopFork(s.SimulationPid)
	if err != nil {
		reply := tgbotapi.NewMessage(message.Chat.ID, "Failed to stop fork: "+err.Error())
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	} else {
		user.SetUserSimulation(strconv.Itoa(int(message.From.ID)), 0, 0)
		reply := tgbotapi.NewMessage(message.Chat.ID, "Fork stopped.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
}

type GetRpcCmd struct{}

func (GetRpcCmd) Command() string {
	return "getrpc"
}

func (GetRpcCmd) HandleCommand(message tgbotapi.Message) {
	if !auth.IsAuth(strconv.Itoa(int(message.From.ID))) {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not authorized to use this command.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	s := user.GetUserStatus(strconv.Itoa(int(message.From.ID)))
	if s.SimulationPid == 0 {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are not running a fork.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	if s.OnlineMode {
		reply := tgbotapi.NewMessage(message.Chat.ID, "You are running online mode.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
		return
	}
	baseUri, _ := os.LookupEnv("TAROLY_WEB_URL")
	json, _ := json2.Marshal(model.EndPoint{
		Uid:  strconv.Itoa(int(message.From.ID)),
		Port: strconv.Itoa(s.SimulationPort),
	})
	sig := hmac.SignWithNonce(string(json))
	reply := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Your rpc endpoint is %srpc/%s", baseUri, sig))
	reply.ReplyToMessageID = message.MessageID
	_, _ = bot.GetBot().Send(reply)
}
