package manager

import (
	"github.com/coutvv/energybot/internal/energy/db"
	"github.com/coutvv/energybot/internal/energy/db/entity"
	"log"
)

// Mediator between interface (telegram) and business logic
type Manager struct {
	Repository db.Repository
}

func (man *Manager) JoinUser(chatId int64, userData entity.User) bool {
	user, err := man.Repository.GetUser(userData.TeleId)
	if err != nil {
		log.Println(err.Error())
		// create user
		user := man.Repository.SaveUser(userData)
		if !user {
			return false
		}
	}
	game, err := man.Repository.GetUnfinishedGame(chatId)
	if err != nil {
		return false
	}
	_, err = man.Repository.JoinGame(user.Id, game)
	if err != nil {
		return false
	}
	return true
}
