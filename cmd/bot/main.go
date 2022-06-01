package main

import (
	"fmt"
	"log"
	"os"

	"github.com/coutvv/energybot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var productService = product.NewService()

func main() {
	godotenv.Load()
	token := os.Getenv("TGTOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			message := update.Message
			log.Printf("[%s] %s", message.From.UserName, message.Text)

			switch message.Command() {
			case "help":
				helpCommand(bot, message)
			case "list":
				listCommand(bot, message)
			case "registry":
				registry(bot, message)
			case "map":
				mapCommand(bot, message)
			case "status":
				status(bot, message)
			default:
				defaultBehavior(bot, message)
			}
		}
	}
}

func listCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	products := "Here all products\n\n"
	for _, product := range productService.List() {
		products += product.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, products)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("You have wrote: %s", inputMsg.Text))
	msg.ReplyToMessageID = inputMsg.MessageID
	bot.Send(msg)
}
