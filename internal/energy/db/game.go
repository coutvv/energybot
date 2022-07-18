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

	// TODO: deck
}

type GamePlayer struct {
	Id     int64
	GameId int64
	UserId int64
	// Some other game properties
	Money int
}

type StationCard struct { // CARD
	Id            int64
	Number        int
	Type          StationType
	Domiki        int
	ResourceCount int
}

type StationType int

const (
	COAL StationType = iota + 1
	OIL
	BURNED // COAL or OIL
	GARBAGE
	NUCLEAR
	GREEN
	STAGE3 // Карта этапа 3
)
