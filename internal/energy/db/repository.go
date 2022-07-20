package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
)

type Repository interface {
	UserRepository
	GameRepository
	PlayerRepository
	Close()
}

type SqliteRepository struct {
	db sql.DB
}

func (sqlRep *SqliteRepository) Close() {
	sqlRep.db.Close()
}

func NewSqliteRepository() Repository {
	return NewSqliteRepositoryCustom("trash/energy-web.db", "./sqlite/create_scheme.sql")
}

func NewSqliteRepositoryCustom(dbFilename string, migrationFile string) Repository {
	var result = SqliteRepository{
		db: newSqliteDb(dbFilename),
	}
	migrationScripts(&result.db, migrationFile)
	return &result
}

func migrationScripts(db *sql.DB, filename string) {
	var buf, err = ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	var script = string(buf)
	_, err = db.Exec(script)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("db migration is completed")
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
