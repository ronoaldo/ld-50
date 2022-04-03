package main

import (
	"errors"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/ronoaldo/ld-50/assets"
)

// Game Options
var (
	width  int = 1920
	height int = 1080
)

// GameScreen indicates the current displayed game screen.
type GameScreen int

var (
	GameScreenTitle     GameScreen = 0
	GameScreenInventory GameScreen = 1
)

var (
	gameExitError = errors.New("user exited the game. what a looser")
)

type Game struct {
	screen      GameScreen
	tickCounter int

	audioContext *audio.Context
	audioPlayer  *audio.Player

	player   *Player
	entities []*Entity
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
	g.player = NewPlayer(g)

	return
}

func (g *Game) Update() error {
	// TODO(ronoaldo): overflow??
	g.tickCounter++

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		log.Printf("Mouse pressed at (%v,%v)", x, y)
	}

	// ESC
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if g.screen == GameScreenInventory {
			g.screen = GameScreenTitle
		} else {
			return gameExitError
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if g.screen == GameScreenTitle {
			g.screen = GameScreenInventory
		}
	}

	for _, e := range g.entities {
		e.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.screen {
	case GameScreenTitle:
		g.TitleScreen(screen)
	case GameScreenInventory:
		g.InventoryScreen(screen)
	}

	for _, e := range g.entities {
		e.Draw(screen)
	}
}

func (g *Game) TitleScreen(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{
		Filter: ebiten.FilterLinear,
	}
	screen.DrawImage(assets.Title, op)
}

func (g *Game) InventoryScreen(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{
		Filter: ebiten.FilterLinear,
	}
	screen.DrawImage(assets.InventoryBackground, op)

	// TODO: convert hard-coded values into constants
	x, y := 39, 189
	for _, droid := range g.player.inv.droids {
		droid.e.x, droid.e.y = float64(x), float64(y)
		droid.e.invisible = false
		x += 192 + 15 // space width + offset
	}

	x, y = 83, 866
	for _, chip := range g.player.inv.chips {
		chip.e.x, chip.e.y = float64(x), float64(y)
		chip.e.invisible = false
		x += 192 + 30 // chip slot width + offset
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Droid Battles")
	ebiten.SetWindowIcon([]image.Image{assets.BlueL1})
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowClosingHandled(false)
	if err := ebiten.RunGame(game); err != nil && err != gameExitError {
		log.Fatal(err)
	}
}
