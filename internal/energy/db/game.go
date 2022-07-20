package db

import (
	"errors"
	"fmt"
	"github.com/coutvv/energybot/internal/energy/db/entity"
	"log"
	"strconv"
	"strings"
)

type GameRepository interface {
	HasStartedGame(chatId int64) bool
	CreateGame(game entity.Game) int64
	GetUnfinishedGame(chatId int64) (entity.Game, error)
	JoinGame(userId int64, game entity.Game) (entity.Player, error)
	ChangeGameState(gameId int64, state entity.State) error
	SaveGameStatus(game entity.Game)
}

func (sqlRep *SqliteRepository) GetUnfinishedGame(chatId int64) (entity.Game, error) {

	row, err := sqlRep.db.Query(
		"SELECT id, chat_id, state, station_market, deck FROM game "+
			"WHERE game.state <> ? AND game.state <> ? AND game.chat_id= ? LIMIT 1;", entity.FINISHED, entity.STOPPED, chatId)
	defer row.Close()
	if err != nil {
		log.Fatal("meh some error", err) // TODO: fix it
	}
	for row.Next() {
		var result = entity.Game{}
		var stationMarket string
		var deck string
		row.Scan(&result.Id, &result.ChatId, &result.Status, &stationMarket, &deck)
		result.Deck = deserializeString(deck)
		result.StationMarket = deserializeString(stationMarket)
		return result, nil
	}
	return entity.Game{}, errors.New("Not found entity")
}

func (sqlRep *SqliteRepository) HasStartedGame(chatId int64) bool {
	row, err := sqlRep.db.Query(
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
	res, err := result.LastInsertId()
	if err != nil {
		log.Fatal("can't get id after inserting game", err)
	}
	game.Id = res
	return res
}

func (sqlRep *SqliteRepository) JoinGame(userId int64, game entity.Game) (entity.Player, error) {
	script := `
		INSERT INTO player (user_id, game_id, money)
		VALUES (?, ?, 0);
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

func (sqlRep *SqliteRepository) SaveGameStatus(game entity.Game) {
	script := `
		UPDATE game
		SET station_market = ?, deck = ?
		WHERE id = ?
	`
	market := serializeArray(game.StationMarket)
	deck := serializeArray(game.Deck)
	sqlRep.db.Exec(script, market, deck, game.Id)
}

func serializeArray(array []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]")
}

func deserializeString(origin string) []int {
	if len(origin) == 0 {
		return []int{}
	} else {
		strValue := strings.Split(origin, ",")
		ary := make([]int, len(strValue))
		for i, value := range strValue {
			ary[i], _ = strconv.Atoi(value)
		}
		return ary
	}
}
