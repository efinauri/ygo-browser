package models

type CardType int

const (
	Monster CardType = 1 << iota
	Spell
	Trap
	Normal
	_
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
	_
	_
	_
	Flip = 1 << (iota + 6)
	Toon
	Xyz
	Pendulum
	QuickPlay
	Continuous
	Equip
	Field
	Counter
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
	QuickPlay:  "QuickPlay",
	Continuous: "Continuous",
	Equip:      "Equip",
	Field:      "Field",
	Counter:    "Counter",
}

func DecodeCardTypes(mask int) []string {
	var types []string
	for bit := Monster; bit <= Counter; bit <<= 1 {
		if mask&int(bit) != 0 {
			types = append(types, cardTypeMap[bit])
		}
	}
	return types
}
