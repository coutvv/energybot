package tg

import (
	"fmt"

	manager2 "github.com/coutvv/energybot/internal/energy/manager"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandHelper struct {
	Bot     *tgbotapi.BotAPI
	Manager manager2.Manager
}

func NewCommandHelper(bot *tgbotapi.BotAPI, manager manager2.Manager) CommandHelper {
	return CommandHelper{
		Bot: bot, Manager: manager,
	}
}

func (ch *CommandHelper) InvokeCommand(message *tgbotapi.Message) {

	switch message.Command() {

	case "help":
		ch.helpCommand(message)
	case "create":
		ch.create(message)
	case "join":
		ch.join(message)
	case "start":
		ch.start(message)
	case "stop":
		ch.finish(message)

	case "phase":
		ch.phase(message)
	case "map":
		ch.mapka(message)
	case "resources":
		ch.resources(message)
	case "market":
		ch.market(message)
	case "status":
		ch.status(message)
	case "moneynote":
		ch.moneynote(message)
	case "skip":
		ch.skip(message)

	default:
		ch.defaultBehavior(message)
	}
}

func (ch *CommandHelper) helpCommand(inputMsg *tgbotapi.Message) {
	helpText := `
	Ну типа игра "Энергосеть"
	Команды:
	Управление игрой
		/create - создать игру (только одна игра на чат)
		/join - присоединиться к созданной игре, пока она не началась (потом нельзя)
		/start - начать игру
		/stop - остановить игру

	Доступные всегда
		/phase - показать фазу игры
		/map - показать карту
		/resources - показать рынок ресурсов
		/market	- электростанции на продаже
		/status {gamer} - состояние игрока (деньги, электростанции)
		/moneynote - показать подсказку по тому сколько денег за домики дадут
		/skip - пропустить ход

	Фаза торгов:
		/choose {1-4-6} - выставить одну из электростанций на торги
		/pass - спасовать
		/doubt 13 - предложить купить за 13
		/up 1 - предложить на 1 доллар дороже
		/replace {1-3} - заменить купленную станцию уже существующей 
	Фаза ресов:
		/buy {coal|oil|trash|nuclear} {count} {station-number} - купить ресурсов на свою станцию 
	Фаза домиков:
		/setup {city-name} - купить домик в городе
	Фаза бюрократии:
		/charge {station-number} - запитать электросеть станцей
	`
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, helpText)
	ch.Bot.Send(msg)
}

func (ch *CommandHelper) defaultBehavior(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID, fmt.Sprintf("Don't know command. You have wrote: %s. Please use /help", inputMsg.Text))
	msg.ReplyToMessageID = inputMsg.MessageID
	ch.Bot.Send(msg)
}

func (ch *CommandHelper) sendMessage(chatId int64, message string) {
	msg := tgbotapi.NewMessage(chatId, message)
	ch.Bot.Send(msg)
}
