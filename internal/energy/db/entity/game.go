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
	StationMarket []int // ids to station cards
	Deck          []int // ids to station cards
}

func (game *Game) TopCardIsSmall() bool {
	var topIsSmall bool
	if len(game.Deck) == 0 {
		topIsSmall = false
	} else {
		topIsSmall = game.Deck[0] < 16
	}
	return topIsSmall
}
