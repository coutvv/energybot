package test

import (
	"github.com/coutvv/energybot/internal/energy/db"
	"github.com/coutvv/energybot/internal/energy/db/entity"
	manager2 "github.com/coutvv/energybot/internal/energy/manager"
	"testing"
)

const testDbName = "../trash/test.db"
const migrationScripts = "../sqlite/create_scheme.sql"

func TestManagerFullGameLifeCycle(t *testing.T) {
	const chatId = 123
	var repository = db.NewSqliteRepositoryCustom(testDbName, migrationScripts)
	var manager = manager2.Manager{Repository: repository}
	manager.CreateGame(chatId)
	manager.JoinUser(chatId, entity.User{TeleId: 2, UserName: "Test"})
	manager.StartGame(chatId)
	manager.FinishGame(chatId)
}
