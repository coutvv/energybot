package tg

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (ch *CommandHelper) phase(inputMsg *tgbotapi.Message) {
	chatId := inputMsg.Chat.ID
	game, _ := ch.Manager.Repository.GetUnfinishedGame(chatId)

	ch.sendMessage(chatId, "Current phase is: "+fmt.Sprint(game.GamePhase))
}
