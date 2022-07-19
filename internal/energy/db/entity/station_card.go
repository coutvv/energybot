package entity

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
