package energy

type EnergyGame struct {
	players map[string]Player
	gameMap GameMap
	phase   Phase
}

func NewGame() EnergyGame {
	return EnergyGame{
		players: make(map[string]Player),
		gameMap: NewSimpleGameMap(),
		phase:   TORG,
	}
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

type Phase int

const (
	TORG Phase = iota
	RESOURCES
	DOMIKI
	MONEY
)
