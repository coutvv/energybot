package db

import (
	"database/sql"
	"log"
)

type User struct {
	Id        int64
	TeleId    int64
	UserName  string
	FirstName string
	LastName  string
	// TODO: don't forget for fucking isBot checking!
}

func CreateUserTable(db *sql.DB) {
	createSQL := `
		CREATE TABLE  IF NOT EXISTS user (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"tele_id" integer UNIQUE,
			"user_name" TEXT,
			"first_name" TEXT,
			"last_name" TEXT
		);
	`

	log.Println("creating user table...")
	stat, err := db.Prepare(createSQL)
	if err != nil {
		log.Fatal(err.Error()) // TODO: no fatal pls
	}
	stat.Exec()
	log.Println("...table user created successfully")
}

func (user *User) create(db *sql.DB) {
	script := `
		INSERT INTO user (tele_id, user_name, first_name, last_name)
		VALUES (?, ?, ?, ?);
	`
	stat, err := db.Prepare(script)
	if err != nil {
		log.Fatal(err.Error()) // TODO: fix fatal
	}
	_, err2 := stat.Exec(user.TeleId, user.UserName, user.FirstName, user.LastName)
	if err2 != nil {
		log.Fatal(err.Error()) // TODO: fix fatal
	}
	log.Println("successing of creating user: " + user.UserName)
}

func userFromDb(row *sql.Rows) User {
	var (
		id        int64
		teleId    int64
		userName  string
		firstName string
		lastName  string
	)
	row.Scan(&id, &teleId, &userName, &firstName, &lastName)
	return User{
		Id:        id,
		TeleId:    teleId,
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
	}
}
