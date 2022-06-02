package phase

type Phase int

const (
	TORG Phase = iota
	RESOURCES
	DOMIKI
	MONEY
)

type PhaseState interface {
	NextPhase() PhaseState
}

type TorgPhase struct {
}

func (tp *TorgPhase) NextPhase() PhaseState {
	return &ResourcesPhase{}
}

type ResourcesPhase struct{}

func (tp *ResourcesPhase) NextPhase() PhaseState {
	return &DomikiPhase{}
}

type DomikiPhase struct{}

func (tp *DomikiPhase) NextPhase() PhaseState {
	return &MoneyPhase{}
}

type MoneyPhase struct{}

func (tp *MoneyPhase) NextPhase() PhaseState {
	return &TorgPhase{}
}
