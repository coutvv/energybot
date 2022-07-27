package manager

import (
	"errors"
	"github.com/coutvv/energybot/internal/energy/db/entity"
	"math/rand"
	"time"
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
		players := man.Repository.GetGamePlayers(game.Id)

		man.prepareMoney(game)
		man.prepareDeck(&game, len(players))
		man.prepareResources(&game)
		man.prepareMapSettings(&game, players)

		man.Repository.ChangeGameState(game.Id, entity.STARTED)
		return nil
	} else {
		return errors.New("game is not in preparing state, so can't start")
	}
}

func (man *Manager) prepareMoney(game entity.Game) {
	players := man.Repository.GetGamePlayers(game.Id)
	for _, player := range players {
		player.Money = 50
		man.Repository.SaveGamePlayerState(player)
	}
}

func (man *Manager) prepareResources(game *entity.Game) {
	game.Coal = 24
	game.Oil = 18
	game.Garbage = 6
	game.Nuclear = 2
	man.Repository.SaveGame(*game)
}

func (man *Manager) FinishGame(chatId int64) error {
	game, err := man.Repository.GetUnfinishedGame(chatId)
	if err != nil {
		return errors.New("something got wrong in db")
	}
	man.Repository.ChangeGameState(game.Id, entity.STOPPED)
	return nil
}

var mapOfRemovingSmallCards = map[int]int{2: 1, 3: 2, 4: 1, 5: 0, 6: 0}
var mapOfRemovingBigCards = map[int]int{2: 5, 3: 6, 4: 3, 5: 0, 6: 0}

func (man *Manager) prepareDeck(game *entity.Game, numOfPlayers int) {
	var smallStations []int
	for i := 3; i < 16; i++ {
		_, err := man.Repository.GetStation(i)
		if err == nil {
			smallStations = append(smallStations, i)
		}
	}
	smallStations = shuffleSlice(smallStations)
	game.StationMarket = smallStations[:8]
	topCard := smallStations[8]
	smallStations = smallStations[9+mapOfRemovingSmallCards[numOfPlayers]:]

	var bigStations []int
	for i := 16; i <= 50; i++ {
		_, err := man.Repository.GetStation(i)
		if err == nil {
			bigStations = append(bigStations, i)
		}
	}
	bigStations = shuffleSlice(bigStations)
	bigStations = bigStations[mapOfRemovingBigCards[numOfPlayers]:]

	deck := append(smallStations, bigStations...)
	deck = shuffleSlice(deck)
	deck = append([]int{topCard}, deck...)
	game.Deck = append(deck, 1000)
	man.Repository.SaveGame(*game)
}

func shuffleSlice(ids []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})
	dest := make([]int, len(ids))
	perm := rand.Perm(len(ids)) // more stupid randoming
	for i, v := range perm {
		dest[v] = ids[i]
	}
	return dest
}
