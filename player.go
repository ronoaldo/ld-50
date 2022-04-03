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
	})
	p.inv.AddDroid(&Droid{
		Name:   "Blue Lvl 2",
		Level:  2,
		Sprite: assets.BlueL2,
	})
	p.inv.AddChip(NewChip())
	p.inv.AddChip(NewChip())
	p.inv.AddChip(NewChip())
	p.inv.AddChip(NewChip())
	p.inv.AddChip(NewChip())
	return p
}
