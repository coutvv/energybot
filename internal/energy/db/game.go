package db

type State int

const (
	PREPARING State = iota + 0
	STARTED
	FINISHED
	STOPPED
)

type Game struct {
	Id     int64
	ChatId int64
	Status State
}

type GamePlayer struct {
	Id     int64
	GameId int64
	UserId int64
	// Some other game properties
}
