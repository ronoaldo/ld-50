package main

import "github.com/hajimehoshi/ebiten/v2"

type Stats struct {
	HP      int
	Strengh int
	Speed   int
}

type Droid struct {
	Name   string
	Level  int
	Sprite *ebiten.Image

	BaseStats Stats

	Chips [6]Chip

	e *Entity
}

type Chip struct {
	ID int

	StatModifier Stats
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
