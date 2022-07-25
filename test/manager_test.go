package test

import (
	"errors"
	"github.com/coutvv/energybot/internal/energy/db"
	"github.com/coutvv/energybot/internal/energy/db/entity"
	manager2 "github.com/coutvv/energybot/internal/energy/manager"
	"log"
	"testing"
)

const testDbName = "../trash/test.db"

var migrationScripts = []string{"../sqlite/create_scheme.sql", "../sqlite/create_cards.sql"}

func TestManagerFullGameLifeCycle(t *testing.T) {

	const chatId = 123
	var repository = db.NewSqliteRepositoryCustom(testDbName, migrationScripts)

	var manager = manager2.Manager{Repository: repository}
	manager.FinishGame(chatId)

	manager.CreateGame(chatId)
	manager.JoinUser(chatId, &entity.User{TeleId: 2, UserName: "Test"})
	manager.JoinUser(chatId, &entity.User{TeleId: 3, UserName: "Bob"})
	startingErr := manager.StartGame(chatId)
	defer manager.FinishGame(chatId)

	if startingErr != nil {
		log.Fatal("something get wrong with starting", startingErr)
	}
	game, _ := repository.GetUnfinishedGame(chatId)
	players := repository.GetGamePlayers(game.Id)
	for _, player := range players {
		if player.Money != 50 {
			panic(errors.New("not setted money"))
		}
	}
	if game.Nuclear != 2 {
		panic("resources doesn't setup")
	}
	if len(game.Deck) == 0 {
		panic("no setup deck")
	}
}
