package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ronoaldo/ld50/assets"
)

var (
	width  int = 640
	height int = 480
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(assets.BlueL1, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Droid Battles")
	ebiten.SetWindowIcon([]image.Image{assets.BlueL1})
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
