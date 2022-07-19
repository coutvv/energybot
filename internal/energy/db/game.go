package db

import (
	"errors"
	"github.com/coutvv/energybot/internal/energy/db/entity"
	"log"
)

type GameRepository interface {
	HasStartedGame(chatId int64) bool
	CreateGame(game entity.Game) int64
	GetUnfinishedGame(chatId int64) (entity.Game, error)
	JoinGame(userId int64, game entity.Game) (entity.Player, error)
	ChangeGameState(gameId int64, state entity.State) error
}

func (sqlRep *SqliteRepository) GetUnfinishedGame(chatId int64) (entity.Game, error) {

	row, err := sqlRep.db.Query(
		"SELECT id, chat_id, state FROM game "+
			"WHERE game.state <> ? AND game.state <> ? AND game.chat_id= ? LIMIT 1;", entity.FINISHED, entity.STOPPED, chatId)
	defer row.Close()
	if err != nil {
		log.Fatal("meh some error", err) // TODO: fix it
	}
	for row.Next() {
		var result = entity.Game{}
		row.Scan(&result.Id, &result.ChatId, &result.Status)
		return result, nil
	}
	return entity.Game{}, errors.New("Not found entity")
}

func (sqRep *SqliteRepository) HasStartedGame(chatId int64) bool {
	row, err := sqRep.db.Query(
		"SELECT * FROM game WHERE game.chat_id = ? AND game.state <> ? AND game.state <> ?;",
		chatId, entity.FINISHED, entity.STOPPED)
	defer row.Close()
	if err != nil {
		log.Fatal(err.Error()) // TODO: mb it should not?
	}
	for row.Next() {
		return true
	}
	return false
}

func (sqlRep *SqliteRepository) CreateGame(game entity.Game) int64 {
	script := `
		INSERT INTO game (chat_id, state)
		VALUES (?, ?);
	`
	stat, err := sqlRep.db.Prepare(script)
	defer stat.Close()
	if err != nil {
		log.Fatal(err.Error()) // TODO: fix fatal
	}
	result, err := stat.Exec(game.ChatId, game.Status)
	if err != nil {
		log.Fatal(err.Error()) // TODO: fix fatal
	}
	println(result)
	return 0 // TODO: fix to from db
}

func (sqlRep *SqliteRepository) JoinGame(userId int64, game entity.Game) (entity.Player, error) {
	script := `
		INSERT INTO game_player (user_id, game_id)
		VALUES (?, ?);
	`
	stat, err := sqlRep.db.Prepare(script)
	defer stat.Close()
	if err != nil {
		log.Fatal(err.Error()) // TODO: fix fatal
	}
	result, err := stat.Exec(userId, game.Id)
	if err != nil {
		log.Println(err.Error())
		return entity.Player{}, errors.New("Can't add to game")
	}
	gamePlayerId, _ := result.LastInsertId()
	return entity.Player{ // TODO: fix to from db
		Id:     gamePlayerId,
		GameId: game.Id,
		UserId: userId,
	}, nil
}

func (sqlRep *SqliteRepository) ChangeGameState(gameId int64, state entity.State) error {
	script := `
		UPDATE game
		SET state = ?
		WHERE id = ?
	`

	stat, err := sqlRep.db.Prepare(script)
	defer stat.Close()
	if err != nil {
		log.Fatal(err.Error()) // TODO: fix fatal
	}
	result, err := stat.Exec(state, gameId)
	if err != nil {
		log.Fatal(err.Error()) // TODO: fix fatal
	}
	println(result)
	return nil
}
