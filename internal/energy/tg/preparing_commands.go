package tg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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

}
func (ch *CommandHelper) start()  {}
func (ch *CommandHelper) finish() {}
