package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

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
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			default:
				defaultBehavior(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Some description of help")
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("You have wrote: %s", inputMsg.Text))
	msg.ReplyToMessageID = inputMsg.MessageID
	bot.Send(msg)
}
