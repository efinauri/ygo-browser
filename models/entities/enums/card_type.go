package enums

type CardType int

const (
	Monster CardType = 1 << iota
	Spell
	Trap
	_
	Normal
	Effect
	Fusion
	Ritual
	_
	Spirit
	Union
	Gemini
	Tuner
	Synchro
	_
	_
	QuickPlay
	Continuous
	Equip
	Field
	Counter
	Flip
	Toon
	Xyz
	Pendulum
	//SpecialSummon// present in theory, unsure if it's useful to show
)

var cardTypeMap = map[CardType]string{
	Monster:    "Monster",
	Spell:      "Spell",
	Trap:       "Trap",
	Normal:     "Normal",
	Effect:     "Effect",
	Fusion:     "Fusion",
	Ritual:     "Ritual",
	Spirit:     "Spirit",
	Union:      "Union",
	Gemini:     "Gemini",
	Tuner:      "Tuner",
	Synchro:    "Synchro",
	Flip:       "Flip",
	Toon:       "Toon",
	Xyz:        "Xyz",
	Pendulum:   "Pendulum",
	QuickPlay:  "Quick-Play",
	Continuous: "Continuous",
	Equip:      "Equip",
	Field:      "Field",
	Counter:    "Counter",
	//SpecialSummon: "Special Summon",
}

func DecodeCardTypes(mask int) []string {
	var types []string
	// do this small iteration first to make sure that a card's main type appears in front
	for bit := Monster; bit <= Trap; bit <<= 1 {
		if mask&int(bit) != 0 {
			types = append(types, cardTypeMap[bit])
			break
		}
	}
	// next iteration is done in reverse, because it looks like the bitmask is ordered to match the card layout in reverse
	for bit := Counter; bit > Trap; bit >>= 1 {
		if mask&int(bit) != 0 {
			types = append(types, cardTypeMap[bit])
		}
	}
	return types
}
