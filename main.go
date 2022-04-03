package main

import (
	"errors"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/ronoaldo/ld-50/assets"
)

// Game Options
var (
	width  int = 1280
	height int = 720
)

type GameScreen int

var (
	TitleScreen GameScreen = 0
)

var (
	gameExitError = errors.New("user exited the game. what a looser")
)

type Game struct {
	screen GameScreen

	audioContext *audio.Context
	audioPlayer  *audio.Player
}

func NewGame() (g *Game, err error) {

	g = &Game{}
	g.audioContext = audio.NewContext(assets.SampleRate)
	g.audioPlayer, err = g.audioContext.NewPlayer(assets.BackgroundMusic)
	if err != nil {
		return nil, err
	}
	g.audioPlayer.SetVolume(0.3)
	g.audioPlayer.Play()
	return
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return gameExitError
	}
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
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Droid Battles")
	ebiten.SetWindowIcon([]image.Image{assets.BlueL1})
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowClosingHandled(false)
	if err := ebiten.RunGame(game); err != nil && err != gameExitError {
		log.Fatal(err)
	}
}
