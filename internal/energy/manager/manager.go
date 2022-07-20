package manager

import (
	"errors"
	"github.com/coutvv/energybot/internal/energy/db"
	"github.com/coutvv/energybot/internal/energy/db/entity"
	"log"
)

// Mediator between interface (telegram) and business logic
type Manager struct {
	Repository db.Repository
}

func (man *Manager) CreateGame(chatId int64) bool {
	// check no started game with chatId or else throw error
	hasStartedGame := man.Repository.HasStartedGame(chatId)

	if hasStartedGame {
		// TODO: return message that can't create game
		return false
	} else {
		man.Repository.CreateGame(entity.Game{
			Status: entity.PREPARING,
			ChatId: chatId,
		})
		return true
	}
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

func (man *Manager) StartGame(chatId int64) error {
	game, err := man.Repository.GetUnfinishedGame(chatId)
	if err != nil {
		return errors.New("not found game")
	}
	if game.Status == entity.PREPARING {
		// prepare player state (add money)

		// prepare deck and station market
		// prepare resources market
		// prepare map
		man.Repository.ChangeGameState(game.Id, entity.STARTED)
		return nil
	} else {
		return errors.New("game is not in preparing state, so can't start")
	}
}

func (man *Manager) FinishGame(chatId int64) error {
	game, err := man.Repository.GetUnfinishedGame(chatId)
	if err != nil {
		return errors.New("something got wrong in db")
	}
	man.Repository.ChangeGameState(game.Id, entity.STOPPED)
	return nil
}
