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
		Name:   "Blue",
		Level:  1,
		Sprite: assets.BlueL1,
	})
	p.inv.AddDroid(&Droid{
		Name:   "Blue",
		Level:  1,
		Sprite: assets.BlueL1,
	})
	return p
}
