package enums

type CardRace int

const (
	Warrior CardRace = 1 << iota
	Spellcaster
	Fairy
	Fiend
	Zombie
	Machine
	Aqua
	Pyro
	Rock
	WingedBeast
	Plant
	Insect
	Thunder
	Dragon
	Beast
	BeastWarrior
	Dinosaur
	Fish
	SeaSerpent
	Reptile
	Psychic
	DivineBeast
	CreatorGod
	Wyrm
)

var cardRaceMap = map[CardRace]string{
	Warrior:      "Warrior",
	Spellcaster:  "Spellcaster",
	Fairy:        "Fairy",
	Fiend:        "Fiend",
	Zombie:       "Zombie",
	Machine:      "Machine",
	Aqua:         "Aqua",
	Pyro:         "Pyro",
	Rock:         "Rock",
	WingedBeast:  "WingedBeast",
	Plant:        "Plant",
	Insect:       "Insect",
	Thunder:      "Thunder",
	Dragon:       "Dragon",
	Beast:        "Beast",
	BeastWarrior: "BeastWarrior",
	Dinosaur:     "Dinosaur",
	Fish:         "Fish",
	SeaSerpent:   "SeaSerpent",
	Reptile:      "Reptile",
	Psychic:      "Psychic",
	DivineBeast:  "DivineBeast",
	CreatorGod:   "CreatorGod",
	Wyrm:         "Wyrm",
}

func DecodeCardRace(mask int) string {
	for bit := Warrior; bit <= Wyrm; bit <<= 1 {
		if mask < int(bit) {
			break
		}
		if mask&int(bit) != 0 {
			return cardRaceMap[bit]
		}
	}
	return ""
}
