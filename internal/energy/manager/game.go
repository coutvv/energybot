package manager

import (
	"errors"
	"github.com/coutvv/energybot/internal/energy/db/entity"
)

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

func (man *Manager) StartGame(chatId int64) error {
	game, err := man.Repository.GetUnfinishedGame(chatId)
	if err != nil {
		return errors.New("not found game")
	}
	if game.Status == entity.PREPARING {
		// prepare player state (add money)
		man.PrepareMoney(game)
		// prepare deck and station market
		// prepare resources market
		// prepare map
		man.Repository.ChangeGameState(game.Id, entity.STARTED)
		return nil
	} else {
		return errors.New("game is not in preparing state, so can't start")
	}
}

func (man *Manager) PrepareMoney(game entity.Game) {
	players := man.Repository.GetGamePlayers(game.Id)
	for _, player := range players {
		player.Money = 50
		man.Repository.SaveGamePlayerState(player)
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
