package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ronoaldo/ld-50/assets"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Stats struct {
	HP      int
	Strengh int
	Speed   int
}

type Droid struct {
	Name      string
	Level     int
	BaseStats Stats
	Chips     [6]Chip

	Sprite *ebiten.Image
	e      *Entity
}

func (d *Droid) Stats() Stats {
	b := d.BaseStats
	s := Stats{
		HP:      b.HP,
		Strengh: b.Strengh,
		Speed:   b.Speed,
	}

	for _, ch := range d.Chips {
		s.HP += ch.StatModifier.HP
		s.Strengh += ch.StatModifier.Strengh
		s.Speed += ch.StatModifier.Speed
	}

	return s
}

type Chip struct {
	ID           uuid.UUID
	StatModifier Stats

	Sprite *ebiten.Image
	e      *Entity
}

func NewChip() *Chip {
	c := &Chip{}
	c.ID = uuid.New()
	c.StatModifier.HP = rand.Intn(11)
	c.StatModifier.Strengh = rand.Intn(11)
	c.StatModifier.Speed = rand.Intn(11)
	// TODO(ronoaldo): randomize chip sprite
	c.Sprite = assets.ChipSpeed
	return c
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
