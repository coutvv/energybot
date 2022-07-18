package main

import (
	"github.com/coutvv/energybot/internal/energy/db"
	manager2 "github.com/coutvv/energybot/internal/energy/manager"
	"github.com/coutvv/energybot/internal/energy/tg"
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

	var repository = db.NewSqliteRepository()
	defer repository.Close()
	var manager = manager2.Manager{Repository: repository}
	commandHelper := tg.NewCommandHelper(bot, manager)

	for update := range updates {
		if update.Message != nil { // If we got a message
			message := update.Message
			log.Printf("[%s] %s", message.From.UserName, message.Text)

			commandHelper.InvokeCommand(message)
		}
	}
}
