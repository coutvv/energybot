package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Repository interface {
	UserRepository
	GameRepository
	Close()
}

type SqliteRepository struct {
	db sql.DB
}

func (sqlRep *SqliteRepository) Close() {
	sqlRep.db.Close()
}

func NewSqliteRepository() Repository {
	var result = SqliteRepository{
		db: newSqliteDb("energy-web.db"),
	}
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
