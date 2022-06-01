package energy

import "fmt"

type Player struct {
	name    string
	money   int
	station []Station // TODO: max 3
	filials []string
}

func (p *Player) State() string {
	result := p.name + "\n"
	result += fmt.Sprintf("money: %d\n", p.money)
	result += "stations:\n"
	for _, station := range p.station {
		result += "\t" + station.toString() + "\n"
	}
	result += "filials: ["
	for _, filial := range p.filials {
		result += filial + ","
	}
	result += "]"
	return result
}

func NewPlayer(name string) Player {
	return Player{
		name:    name,
		money:   50,
		station: []Station{},
		filials: []string{},
	}
}

type Station struct {
	price    int
	power    int
	resType  Resource
	resCount int
}

func (s *Station) toString() string {
	return fmt.Sprintf("type: %s\tpower: %d\tprice:%d\tresources:%d", s.resType.toString(), s.power, s.price, s.resCount)
}

type Resource int

const (
	Coal Resource = iota
	Oil
	Nuclear
	Trash
	Eco
)

func (r *Resource) toString() string {
	var result string
	if *r == Coal {
		result = "coal"
	}
	if *r == Oil {
		result = "oil"
	}
	if *r == Trash {
		result = "trash"
	}
	if *r == Nuclear {
		result = "nuclear"
	}
	if *r == Eco {
		result = "eco"
	}
	return result
}
