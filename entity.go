package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Entity struct {
	name    string
	texture *ebiten.Image

	counter int

	frameCount   int
	currentFrame int

	x float64
	y float64

	vx float64
	vy float64
}

func NewEntity(name string, texture *ebiten.Image) *Entity {
	e := &Entity{
		name:       name,
		texture:    texture,
		frameCount: texture.Bounds().Max.Y / 64,
		vx:         2,
		vy:         2,
	}
	log.Printf("New entity at %v, %v frames", e.name, e.frameCount)
	return e
}

func (e *Entity) Update() {
	// Move with mouse
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		log.Printf("Mouse pressed at (%v,%v)", x, y)
		e.x = float64(x)
		e.y = float64(y)
	}
	// Move with arrows
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		e.x = e.x - e.vx
		if e.x < 0 {
			e.x = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		e.x = e.x + e.vx
		if e.x > float64(width) {
			e.x = float64(width)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		e.y = e.y - e.vy
		if e.y < 0 {
			e.y = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		e.y = e.y + e.vy
		if e.y > float64(height) {
			e.y = float64(height)
		}
	}

	e.counter++
}

func (e *Entity) Draw(screen *ebiten.Image) {
	geom := ebiten.GeoM{}
	geom.Translate(e.x, e.y)
	geom.Scale(3.0, 3.0)

	op := &ebiten.DrawImageOptions{
		GeoM:   geom,
		Filter: ebiten.FilterNearest,
	}

	// Vertical Sprite
	// [0,0]
	//          [64, 64]
	// [0,64*frame]
	//          [64, 64*(frame+1)]
	framePos := (e.counter / 6) % e.frameCount
	frame := image.Rect(0, 64*framePos, 64, 64*(framePos+1))
	screen.DrawImage(e.texture.SubImage(frame).(*ebiten.Image), op)
}
