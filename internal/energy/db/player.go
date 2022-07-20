package db

import (
	"github.com/coutvv/energybot/internal/energy/db/entity"
	"log"
)

type PlayerRepository interface {
	GetGamePlayers(gameId int64) []entity.Player
	SaveGamePlayerState(player entity.Player)
}

func (sqlRep *SqliteRepository) GetGamePlayers(gameId int64) []entity.Player {
	script := `SELECT * FROM PLAYER where game_id = ?;`
	rows, err := sqlRep.db.Query(script, gameId)
	defer rows.Close()
	if err != nil {
		log.Fatal(err.Error()) // TODO: mb fix this maaan...
	}
	var result []entity.Player
	for rows.Next() {
		player := entity.Player{}
		rows.Scan(&player.Id, &player.UserId, &player.GameId, &player.Money)
		result = append(result, player)
	}
	return result
}

func (sqlRep *SqliteRepository) SaveGamePlayerState(player entity.Player) {
	script := `
		UPDATE player
		SET money = ?
		WHERE id = ?;
	`
	_, err := sqlRep.db.Exec(script, player.Money, player.Id)
	if err != nil {
		log.Fatal(err.Error()) // oh god what I should do now?
	}
}
