package main

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ronoaldo/ld-50/assets"
)

// Unit stats for both player and enemy.
type Stats struct {
	HP      int
	Strengh int
	Speed   int
}

// Chip a stat modifier that can be applied to the player/enemy.
type Chip struct {
	ID           uuid.UUID
	StatModifier Stats

	Sprite *ebiten.Image
	e      *Entity
}

// NewChip creates a new random stat modifier
func NewChip() *Chip {
	c := &Chip{}
	c.ID = uuid.New()

	// Randomize new chip stats player can choose
	c.StatModifier.HP = rand.Intn(11)
	c.StatModifier.Strengh = rand.Intn(11)
	c.StatModifier.Speed = rand.Intn(11)

	// Randomize chip graphics
	if c.StatModifier.HP > c.StatModifier.Speed &&
		c.StatModifier.HP > c.StatModifier.Strengh {
		c.Sprite = assets.ChipLife
	} else if c.StatModifier.Strengh > c.StatModifier.HP &&
		c.StatModifier.Strengh > c.StatModifier.Speed {
		c.Sprite = assets.ChipStrength
	} else {
		c.Sprite = assets.ChipSpeed
	}

	return c
}

// Droid is a base unit type that is a playable character.
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

type BattleUnit struct {
	Name string

	// Final stats from the base + modifiers
	Stats Stats

	// Skills this unit can use
	Skills [3]Skill
}

type SkillEffect func(p, e *BattleUnit)

type Skill struct {
	Name   string
	Effect SkillEffect
}

var (
	AttackSkill = Skill{
		Name: "Attack",
		Effect: func(p, e *BattleUnit) {
			str := p.Stats.Strengh
			e.Stats.HP -= str
		},
	}
)
