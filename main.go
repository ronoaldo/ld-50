package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ronoaldo/ld-50/assets"
)

var (
	width  int = 1280
	height int = 720
)

type GameScreen int

var (
	TitleScreen GameScreen = 0
)

type Game struct {
	screen GameScreen
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.screen {
	case TitleScreen:
		g.TitleScreen(screen)
	}
}

func (g *Game) TitleScreen(screen *ebiten.Image) {
	geom := adaptScale(assets.Title, screen)
	op := &ebiten.DrawImageOptions{
		GeoM:   geom,
		Filter: ebiten.FilterLinear,
	}
	screen.DrawImage(assets.Title, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func adaptScale(src, dst *ebiten.Image) ebiten.GeoM {
	srcd := src.Bounds()
	dstd := dst.Bounds()

	scalingX := float64(dstd.Max.X) / float64(srcd.Max.X)
	scalingY := float64(dstd.Max.Y) / float64(srcd.Max.Y)

	geom := ebiten.GeoM{}
	geom.Scale(float64(scalingX), float64(scalingY))
	return geom
}

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Droid Battles")
	ebiten.SetWindowIcon([]image.Image{assets.BlueL1})
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
