package manager

import (
	"github.com/coutvv/energybot/internal/energy/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Manager struct {
	Repository db.Repository
}

func (man *Manager) RegistryGamer(inputMsg *tgbotapi.Message) (db.User, bool) {
	var teleId = inputMsg.From.ID
	user, err := man.Repository.GetUser(teleId)
	if err != nil {
		var newUser = db.User{
			TeleId:    teleId,
			UserName:  inputMsg.From.UserName,
			FirstName: inputMsg.From.FirstName,
			LastName:  inputMsg.From.LastName,
		}
		man.Repository.SaveUser(newUser)
		return newUser, true
	} else {
		return user, false
	}
}
