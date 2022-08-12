package manager

import (
	"errors"
	"log"

	"github.com/coutvv/energybot/internal/energy/db"
	"github.com/coutvv/energybot/internal/energy/db/entity"
)

// Mediator between interface (telegram) and business logic
type Manager struct {
	Repository db.Repository
}

func (man *Manager) JoinUser(chatId int64, userData *entity.User) bool {
	user, err := man.Repository.GetUser(userData.TeleId)
	if err != nil {
		log.Println(err.Error())
		// create user
		hasCreated := man.Repository.SaveUser(userData)
		user = *userData
		if !hasCreated {
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

func (man *Manager) prepareMapSettings(game *entity.Game, players []entity.Player) error {
	playerCount := len(players)
	switch playerCount {
	case 2:
		game.Regions = []string{"A", "B"}
	case 3:
		game.Regions = []string{"A", "B", "C"}
	case 4:
		game.Regions = []string{"A", "B", "C", "D"}
	case 5:
		game.Regions = []string{"A", "B", "C", "D", "E"}
	case 6:
		game.Regions = []string{"A", "B", "C", "D", "E", "F"}
	default:
		return errors.New("incorrect number of players need more!")
	}
	man.Repository.SaveGame(*game)
	return nil
}

func (man *Manager) prepareGameOrder(game *entity.Game, players []entity.Player) {
	game.GameOrder = game.ComputeGameOrder(players)
	man.Repository.SaveGame(*game)
}
