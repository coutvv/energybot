package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (ch *CommandHelper) create(inputMsg *tgbotapi.Message) {
	created := ch.Manager.CreateGame(inputMsg.Chat.ID)
	var message = "Can't create game in this chat until last one in process..."
	if created {
		message = "Game has created!"
	}
	ch.Bot.Send(
		tgbotapi.NewMessage(inputMsg.Chat.ID, message),
	)
}

func (ch *CommandHelper) join(inputMsg *tgbotapi.Message) {
	userId := inputMsg.From.ID
	chatId := inputMsg.Chat.ID

	if ch.Manager.JoinUser(inputMsg) {
		ch.sendMessage(chatId, "User has joined: "+strconv.FormatInt(userId, 10))
	} else {
		ch.sendMessage(chatId, "User can't join to game "+strconv.FormatInt(userId, 10))
	}

}
func (ch *CommandHelper) start(inputMsg *tgbotapi.Message) {
	chatId := inputMsg.Chat.ID
	err := ch.Manager.StartGame(chatId)
	if err != nil {
		ch.sendMessage(chatId, err.Error())
	} else {
		ch.sendMessage(chatId, "so let the game begin!")
	}
}
func (ch *CommandHelper) finish(inputMsg *tgbotapi.Message) {
	chatId := inputMsg.Chat.ID
	err := ch.Manager.FinishGame(chatId)
	if err != nil {
		ch.sendMessage(chatId, "can't finish game it hasn't create yet or stopped already")
	} else {
		ch.sendMessage(chatId, "game has stopped")
	}
}
