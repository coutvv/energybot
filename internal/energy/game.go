package energy

import (
	"reflect"

	"github.com/coutvv/energybot/internal/energy/phase"
)

type EnergyGame struct {
	players map[string]Player
	gameMap GameMap
	phase   phase.PhaseState
}

func NewGame() EnergyGame {
	return EnergyGame{
		players: make(map[string]Player),
		gameMap: NewSimpleGameMap(),
		phase:   &phase.TorgPhase{},
	}
}

func (eg *EnergyGame) Skip(playerName string) {
	// TODO: fix this
	eg.phase = eg.phase.NextPhase()
}

func (eg *EnergyGame) CurrentPhase() string {
	return reflect.TypeOf(eg.phase).String() // TODO: to normal shit
}

func (eg *EnergyGame) PlayerStats(name string) string {
	player, ok := eg.players[name]
	if ok {
		return player.State()
	} else {
		return "no player with that username " + name
	}
}

func (eg *EnergyGame) RegistryGamer(name string) bool {
	if _, ok := eg.players[name]; ok {
		// TODO: fuck stupid gamer to the fuck
		return false
	}
	eg.players[name] = NewPlayer(name)
	return true
}

func (eg *EnergyGame) MapString() string {
	return eg.gameMap.Show()
}
