package entity

import (
	"errors"
	"log"
)

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

	StationMarket []int // ids to station cards
	Deck          []int // ids to station cards

	Coal    int
	Oil     int
	Garbage int
	Nuclear int

	Regions []string // region ids for this game
}

func (game *Game) BuyResources(player Player, resType StationType, count int) bool {
	cost, err := game.ResourcesCost(resType, count)
	if err != nil {
		return false
	}
	if player.Money-cost < 0 {
		return false
	}
	player.Money -= cost
	switch resType {
	case COAL:
		game.Coal -= count
	case GARBAGE:
		game.Garbage -= count
	case OIL:
		game.Oil -= count
	case NUCLEAR:
		game.Nuclear -= count
	default:
		log.Fatal("oh no, how it is going right like that?")
	}
	return true
}

func (game *Game) ResourcesCost(resType StationType, count int) (int, error) {
	result := 0
	balance := 0
	switch resType {
	case COAL:
		balance = game.Coal
		for i := 0; i < count; i++ {
			balance = balance - 1
			unitPrice := unitPriceByDivide[balance/3]
			result += unitPrice
		}
	case GARBAGE:
		balance = game.Garbage
		for i := 0; i < count; i++ {
			balance = balance - 1
			unitPrice := unitPriceByDivide[balance/3]
			result += unitPrice
		}
	case OIL:
		balance = game.Oil
		for i := 0; i < count; i++ {
			balance = balance - 1
			unitPrice := unitPriceByDivide[balance/3]
			result += unitPrice
		}
	case NUCLEAR:
		balance = game.Nuclear
		for i := 0; i < count; i++ {
			result += nuclearPrice[balance]
			balance--
		}
	default:
		log.Println("No found resource")
		return 0, errors.New("Incorrect type of resources")
	}
	if balance < 0 {
		return 0, errors.New("no resources sir for buying")
	} else {
		return result, nil
	}
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

var nuclearPrice = initNuclearPriceTable()

func initNuclearPriceTable() map[int]int {
	result := map[int]int{}
	for i := 1; i <= 12; i++ {
		if i > 4 {
			result[i] = 13 - i
		} else {
			result[i] = 18 - (i * 2)
		}
	}
	return result
}

var unitPriceByDivide = map[int]int{
	7: 1, 6: 2, 5: 3, 4: 4,
	3: 5, 2: 6, 1: 7, 0: 8,
}
