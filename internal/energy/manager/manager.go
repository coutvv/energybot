package manager

import (
	"github.com/coutvv/energybot/internal/energy/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Mediator between interface (telegram) and business logic
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

func (man *Manager) CreateGame(inputMsg *tgbotapi.Message) bool {
	chatId := inputMsg.Chat.ID
	// check no started game with chatId or else throw error
	hasStartedGame := man.Repository.HasStartedGame(chatId)

	if hasStartedGame {
		// TODO: return message that can't create game
		return false
	} else {
		man.Repository.CreateGame(db.Game{
			Status: db.PREPARING,
			ChatId: chatId,
		})
		return true
	}
}

func (man *Manager) JoinUser(inputMsg *tgbotapi.Message) {

}

func (man *Manager) StartGame(chatId int64) {

}

func (man *Manager) FinishGame(chatId int64) {

}
