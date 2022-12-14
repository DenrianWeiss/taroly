package handlers

import (
	"errors"
	"github.com/DenrianWeiss/taroly/service/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var h = map[string]CommandHandlers{}

type CommandHandlers interface {
	Command() string
	HandleCommand(message tgbotapi.Message)
}

func RegisterCommandHandlers(commandHandlers CommandHandlers) {
	h[commandHandlers.Command()] = commandHandlers
}

func GetCommandHandlers(command string) (CommandHandlers, error) {
	if commandHandlers, ok := h[command]; ok {
		return commandHandlers, nil
	}
	return nil, errors.New("command not found")
}

func RegisterCommands() {
	RegisterCommandHandlers(AuthCmd{})
	RegisterCommandHandlers(IdCmd{})
	// User Status
	RegisterCommandHandlers(OnlineModeCmd{})
	RegisterCommandHandlers(SetChainCmd{})
	RegisterCommandHandlers(GetChainCmd{})
	RegisterCommandHandlers(SetAccountCmd{})
	RegisterCommandHandlers(GetAccountCmd{})
	// Fork Management
	RegisterCommandHandlers(NewForkCmd{})
	RegisterCommandHandlers(StopForkCmd{})
	// Encode/Decode
	RegisterCommandHandlers(Byte4Cmd{})
	RegisterCommandHandlers(DecodeCmd{})
	RegisterCommandHandlers(EncodeCmd{})
	// Read only methods
	RegisterCommandHandlers(GetBalanceCmd{})
	RegisterCommandHandlers(CallCmd{})
	RegisterCommandHandlers(CallDataCmd{})
	// Write methods
	RegisterCommandHandlers(SetBalanceCmd{})
	RegisterCommandHandlers(TransactCmd{})
	//
	RegisterCommandHandlers(SetBalanceCmd{})
	RegisterCommandHandlers(TraceCmd{})
	RegisterCommandHandlers(GetRpcCmd{})
}

func init() {
	RegisterCommands()
}

func DispatchCommand(message tgbotapi.Message) {
	if commandHandlers, err := GetCommandHandlers(message.Command()); err == nil {
		commandHandlers.HandleCommand(message)
	} else {
		if message.Command() == "" {
			return
		}
		reply := tgbotapi.NewMessage(message.Chat.ID, "Command not found.")
		reply.ReplyToMessageID = message.MessageID
		_, _ = bot.GetBot().Send(reply)
	}
}
