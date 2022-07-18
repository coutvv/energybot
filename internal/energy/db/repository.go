package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
)

type Repository interface {
	GetUser(teleId int64) (User, error)
	SaveUser(user User) bool
	Close()
	HasStartedGame(chatId int64) bool
	CreateGame(game Game) int64
	GetUnfinishedGame(chatId int64) (Game, error)
	JoinGame(userId int64, game Game) (GamePlayer, error)
	ChangeGameState(gameId int64, state State) error
}

type SqliteRepository struct {
	db sql.DB
}

func (sqlRep *SqliteRepository) GetUnfinishedGame(chatId int64) (Game, error) {

	row, err := sqlRep.db.Query(
		"SELECT id, chat_id, state FROM game WHERE game.state <> ? AND game.chat_id= ? LIMIT 1;", FINISHED, chatId)
	defer row.Close()
	if err != nil {
		log.Fatal("meh some error", err) // TODO: fix it
	}
	for row.Next() {
		var result = Game{}
		row.Scan(&result.Id, &result.ChatId, &result.Status)
		return result, nil
	}
	return Game{}, errors.New("Not found entity")
}

func NewSqliteRepository() Repository {
	var result = SqliteRepository{
		db: newSqliteDb("energy-web.db"),
	}
	result.Setup()
	return &result
}

func newSqliteDb(dbFilename string) sql.DB {
	if _, err := os.Stat(dbFilename); errors.Is(err, os.ErrNotExist) {
		os.Remove(dbFilename)
		log.Println("creating db file")
		file, err := os.Create(dbFilename)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("db is created!")
	}
	db, _ := sql.Open("sqlite3", "./"+dbFilename)
	log.Println("Connection to db is established")
	return *db
}

func (sqlRep *SqliteRepository) Close() {
	sqlRep.db.Close()
}

func (sqRep SqliteRepository) GetUser(teleId int64) (User, error) {
	row, err := sqRep.db.Query("SELECT * FROM user WHERE user.tele_id = ? LIMIT 1;", teleId)
	defer row.Close()
	if err != nil {
		log.Fatal(err.Error()) // TODO: mb it should not?
	}
	for row.Next() {
		return userFromDb(row), nil
	}
	return User{}, errors.New("no user here man with teleId: " + strconv.FormatInt(teleId, 10))
}

func (sqRep *SqliteRepository) SaveUser(user User) bool {
	script := `
		INSERT INTO user (tele_id, user_name, first_name, last_name)
		VALUES (?, ?, ?, ?);
	`
	stat, err := sqRep.db.Prepare(script)
	defer stat.Close()
	if err != nil {
		log.Fatal(err.Error()) // TODO: fix fatal
		return false
	}
	_, err2 := stat.Exec(user.TeleId, user.UserName, user.FirstName, user.LastName)
	if err2 != nil {
		log.Fatal(err.Error()) // TODO: fix fatal
		return false
	}
	log.Println("successing of creating user: " + user.UserName)
	return true
}

func (sqRep *SqliteRepository) Setup() {
	CreateUserTable(&sqRep.db)
}

func (sqRep *SqliteRepository) HasStartedGame(chatId int64) bool {
	row, err := sqRep.db.Query("SELECT * FROM game WHERE game.chat_id = ?;", chatId)
	defer row.Close()
	if err != nil {
		log.Fatal(err.Error()) // TODO: mb it should not?
	}
	for row.Next() {
		return true
	}
	return false
}

func (sqlRep *SqliteRepository) CreateGame(game Game) int64 {
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

func (sqlRep *SqliteRepository) JoinGame(userId int64, game Game) (GamePlayer, error) {
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
		return GamePlayer{}, errors.New("Can't add to game")
	}
	println("fuck it")
	gamePlayerId, _ := result.LastInsertId()
	println("fuck it")
	return GamePlayer{ // TODO: fix to from db
		Id:     gamePlayerId,
		GameId: game.Id,
		UserId: userId,
	}, nil
}

func (sqlRep *SqliteRepository) ChangeGameState(gameId int64, state State) error {
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
