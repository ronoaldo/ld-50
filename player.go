package main

import "github.com/ronoaldo/ld-50/assets"

type Player struct {
	name string

	game *Game
	inv  *Inventory
}

func NewPlayer(g *Game) *Player {
	p := &Player{
		name: "Bob",
	}
	p.game = g
	p.inv = NewInventory(p)
	p.inv.AddDroid(&Droid{
		Name:   "Blue Lvl 1",
		Level:  1,
		Sprite: assets.BlueL1,
		BaseStats: Stats{
			HP:      100,
			Strengh: 20,
			Speed:   10,
		},
		Skills: [3]Skill{
			BasicAttackSkill,
			HealSkill,
			UltimateAttackSkill,
		},
	})
	p.inv.AddDroid(&Droid{
		Name:   "Octopus",
		Level:  2,
		Sprite: assets.OctopusEnemi,
		BaseStats: Stats{
			HP:      100,
			Strengh: 20,
			Speed:   10,
		},
	})
	p.inv.AddChip(NewChip())
	p.inv.AddChip(NewChip())
	p.inv.AddChip(NewChip())
	p.inv.AddChip(NewChip())
	p.inv.AddChip(NewChip())
	return p
}
