package energy

import "fmt"

type GameMap struct {
	cities   []City
	cables   []CityCable
	occupied map[string][]string // cityName to playerName
}

func NewSimpleGameMap() GameMap {
	cities := [2]City{
		{"kazan"},
		{"almetievsk"},
	}
	return GameMap{
		cities: cities[:],
		cables: []CityCable{
			{cities: cities, Price: 10},
		},
		occupied: make(map[string][]string),
	}
}

type City struct {
	name string
}

type CityCable struct {
	cities [2]City
	Price  int
}

func (cg *CityCable) toString() string {
	return fmt.Sprintf("%s <--%d$--> %s", cg.cities[0], cg.Price, cg.cities[1])
}

func (gm *GameMap) Show() string {
	var result = ""
	for _, cable := range gm.cables {
		result += cable.toString() + "\n"
	}
	return result
}
