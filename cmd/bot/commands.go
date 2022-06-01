package main

import (
	"fmt"

	"github.com/coutvv/energybot/internal/energy"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TODO: may be delete stupid bot API from it (?)
func helpCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	helpText := `
	Ну типа "Энергосеть"
	Команды:
	Доступные всегда
		/registry - войти в игру
		/phase - показать фазу игры
		/map - показать карту
		/resources - показать рынок ресурсов
		/market	- электростанции на продаже
		/status {gamer} - состояние игрока (деньги, электростанции)
		/moneynote - показать подсказку по 
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
	bot.Send(msg)
}

var game = energy.NewGame()

// persistent commands

func registry(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	var name = inputMsg.From.UserName
	var registred = game.RegistryGamer(name)
	var msgText string
	if registred {
		msgText = fmt.Sprintf("new player %s has been registred", name)
	} else {
		msgText = "player with name " + name + " is already playing!"
	}
	bot.Send(
		tgbotapi.NewMessage(inputMsg.Chat.ID, msgText),
	)
}

func phase() {}

func mapCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msgText := game.MapString()
	bot.Send(
		tgbotapi.NewMessage(inputMsg.Chat.ID, msgText),
	)
}

func resources() {}

func market() {}

func status(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msgText := game.PlayerStats(inputMsg.From.UserName)
	bot.Send(
		tgbotapi.NewMessage(inputMsg.Chat.ID, msgText),
	)
}

func moneynote() {}

func skip() {}

// torgi phase

func choose() {}

func pass() {}

func doubt() {}

func up() {}

func replace() {}

// resources phase

func buy() {}

// domiki phase

func setup() {}

// buerucracy phace
func charge() {}
