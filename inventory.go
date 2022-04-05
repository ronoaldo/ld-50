package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Inventory struct {
	player *Player
	droids []*Droid
	chips  []*Chip
}

func NewInventory(p *Player) *Inventory {
	i := &Inventory{
		player: p,
	}
	return i
}

func (i *Inventory) AddDroid(d *Droid) {
	i.droids = append(i.droids, d)
	d.e = NewEntity(d.Name, d.Sprite)
	d.e.invisible = true
	i.player.game.entities = append(i.player.game.entities, d.e)
}

func (i *Inventory) AddChip(c *Chip) {
	i.chips = append(i.chips, c)
	c.e = NewEntity(c.ID.String(), c.Sprite)
	c.e.invisible = true
	i.player.game.entities = append(i.player.game.entities, c.e)
}
