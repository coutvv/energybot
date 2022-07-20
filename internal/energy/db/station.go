package db

import (
	"errors"
	"github.com/coutvv/energybot/internal/energy/db/entity"
	"log"
)

type StationCardRepository interface {
	GetStation(id int) (entity.StationCard, error)
}

func (sqlRep *SqliteRepository) GetStation(id int) (entity.StationCard, error) {
	rows, err := sqlRep.db.Query("SELECT * FROM STATION_CARD WHERE id = ?", id)
	defer rows.Close()
	if err != nil {
		log.Fatal("incorrect id or something " + err.Error())
	}

	for rows.Next() {
		result := entity.StationCard{}
		rows.Scan(&result.Id, &result.Number, &result.Domiki, &result.ResourceCount, &result.Type)
		return result, nil
	}
	return entity.StationCard{}, errors.New("no card")
}
