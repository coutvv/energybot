package tg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (ch *CommandHelper) phase(inputMsg *tgbotapi.Message) {
	chatId := inputMsg.Chat.ID
	game, _ := ch.Manager.Repository.GetUnfinishedGame(chatId)
	var phase string
	switch game.GamePhase {
	case 0:
		phase = "torgi"
	case 1:
		phase = "torgi"
	case 2:
		phase = "resources"
	case 3:
		phase = "domiki"
	case 4:
		phase = "money"
	default:
		log.Fatal("incorrect phase of this game")
	}

	ch.sendMessage(chatId, "Current phase is: "+phase)
}
