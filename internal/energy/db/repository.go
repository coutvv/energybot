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
	StationCardRepository
	Close()
}

type SqliteRepository struct {
	db sql.DB
}

func (sqlRep *SqliteRepository) Close() {
	sqlRep.db.Close()
}

func NewSqliteRepository() Repository {
	return NewSqliteRepositoryCustom("trash/energy-web.db",
		[]string{"./sqlite/create_scheme.sql", "./sqlite/create_cards.sql"})
}

func NewSqliteRepositoryCustom(dbFilename string, migrationFile []string) Repository {
	var result = SqliteRepository{
		db: newSqliteDb(dbFilename),
	}
	for _, filename := range migrationFile {
		result.MigrationScripts(filename)
	}
	return &result
}

func (sqlRep *SqliteRepository) MigrationScripts(filename string) {
	var buf, err = ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	var script = string(buf)
	_, err = sqlRep.db.Exec(script)
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
