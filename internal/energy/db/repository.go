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
}

type SqliteRepository struct {
	db sql.DB
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
