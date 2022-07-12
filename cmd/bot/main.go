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
			message := update.Message
			log.Printf("[%s] %s", message.From.UserName, message.Text)

			switch message.Command() {
			case "help":
				helpCommand(bot, message)
			case "registry":
				registry(bot, message)
			case "map":
				mapCommand(bot, message)
			case "status":
				status(bot, message)
			case "phase":
				phase(bot, message)
			case "skipPhase":
				skip(bot, message)
			case "moneynote":
				moneynote(bot, message)
			default:
				defaultBehavior(bot, message)
			}
		}
	}
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("You have wrote: %s", inputMsg.Text))
	msg.ReplyToMessageID = inputMsg.MessageID
	bot.Send(msg)
}
