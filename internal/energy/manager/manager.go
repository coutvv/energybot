package manager

import (
	"errors"
	"github.com/coutvv/energybot/internal/energy/db"
	"github.com/coutvv/energybot/internal/energy/db/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Mediator between interface (telegram) and business logic
type Manager struct {
	Repository db.Repository
}

func (man *Manager) RegistryGamer(inputMsg *tgbotapi.Message) (entity.User, bool) {
	var teleId = inputMsg.From.ID
	user, err := man.Repository.GetUser(teleId)
	if err != nil {
		var newUser = entity.User{
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

func (man *Manager) JoinUser(inputMsg *tgbotapi.Message) bool {
	teleId := inputMsg.From.ID
	user, err := man.Repository.GetUser(teleId)
	if err != nil {
		// create user
		user = entity.User{TeleId: teleId, UserName: inputMsg.From.UserName, FirstName: inputMsg.From.FirstName, LastName: inputMsg.From.LastName}
		created := man.Repository.SaveUser(user)
		if !created {
			return false
		}
	}
	game, err := man.Repository.GetUnfinishedGame(inputMsg.Chat.ID)
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
