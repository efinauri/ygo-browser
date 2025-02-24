package entities

import (
	"fmt"
	"strings"
	"yugioh-browser/models/entities/enums"
)

type Card struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Types     string `json:"types"`
	Atk       int    `json:"atk"`
	Def       int    `json:"def"`
	Level     int    `json:"level"`
	Race      string `json:"race"`
	Attribute string `json:"attribute"`
}

func (c *Card) Sanitize(typeMask int, raceMask int, attributeMask int) {
	types := enums.DecodeCardTypes(typeMask)
	c.Types = strings.Join(types, "/")
	c.Race = enums.DecodeCardRace(raceMask)
	c.Attribute = enums.DecodeCardAttribute(attributeMask)
}

func (c *Card) LevelStr() string { return fmt.Sprintf("%d", c.Level) }
func (c *Card) AtkStr() string {
	if c.Atk < 0 {
		return "?"
	} else {
		return fmt.Sprintf("%d", c.Atk)
	}
}
func (c *Card) DefStr() string {
	if c.Def < 0 {
		return "?"
	} else {
		return fmt.Sprintf("%d", c.Def)
	}
}
