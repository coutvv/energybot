package db

import (
	"database/sql"
	"errors"
	"github.com/coutvv/energybot/internal/energy/db/entity"
	"log"
	"strconv"
)

type UserRepository interface {
	GetUser(teleId int64) (entity.User, error)
	SaveUser(user entity.User) bool
}

func (sqRep *SqliteRepository) SaveUser(user entity.User) bool {
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

func (sqRep SqliteRepository) GetUser(teleId int64) (entity.User, error) {
	row, err := sqRep.db.Query("SELECT * FROM user WHERE user.tele_id = ? LIMIT 1;", teleId)
	defer row.Close()
	if err != nil {
		log.Fatal(err.Error()) // TODO: mb it should not?
	}
	for row.Next() {
		return userFromDb(row), nil
	}
	return entity.User{}, errors.New("no user here with teleId: " + strconv.FormatInt(teleId, 10))
}

func userFromDb(row *sql.Rows) entity.User {
	var (
		id        int64
		teleId    int64
		userName  string
		firstName string
		lastName  string
	)
	row.Scan(&id, &teleId, &userName, &firstName, &lastName)
	return entity.User{
		Id:        id,
		TeleId:    teleId,
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
	}
}
