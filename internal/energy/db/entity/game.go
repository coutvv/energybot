package entity

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
