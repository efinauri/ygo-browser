package enums

type CardAttribute int

const (
	EARTH CardAttribute = 1 << iota
	WATER
	FIRE
	WIND
	LIGHT
	DARK
	DIVINE
)

var cardAttributeMap = map[CardAttribute]string{
	EARTH:  "EARTH",
	WATER:  "WATER",
	FIRE:   "FIRE",
	WIND:   "WIND",
	LIGHT:  "LIGHT",
	DARK:   "DARK",
	DIVINE: "DIVINE",
}

func DecodeCardAttribute(mask int) string {
	for bit := EARTH; bit <= DIVINE; bit <<= 1 {
		if mask < int(bit) {
			break
		}
		if mask&int(bit) != 0 {
			return cardAttributeMap[bit]
		}
	}
	return ""
}
