package db

import (
	"log"

	"github.com/coutvv/energybot/internal/energy/db/entity"
)

type RegionMapRepository interface {
	GetAllCities() []entity.City
	GetAllCables() []entity.Cable
}

func (sqlRep *SqliteRepository) GetAllCities() []entity.City{
	script := `SELECT * FROM CITY`

	rows, err := sqlRep.db.Query(script)
	if err != nil {
		log.Fatal(err.Error()) // TODO: mb fix this maaan...
	}
	defer rows.Close()
	var result []entity.City
	for rows.Next() {
		city := entity.City{}
		rows.Scan(&city.Id, &city.Name, &city.Region)
		result = append(result, city)
	}
	return result
}

func (sqlRep *SqliteRepository) GetAllCables() []entity.Cable{
	script := `SELECT * FROM CABLE`

	rows, err := sqlRep.db.Query(script)
	if err != nil {
		log.Fatal(err.Error()) 
	}
	defer rows.Close()
	var result []entity.Cable
	for rows.Next() {
		cable := entity.Cable{}
		rows.Scan(&cable.Src, &cable.Dest, &cable.Price)
		result = append(result, cable)
	}

	return result
}