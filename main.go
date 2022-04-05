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
	GameScreenBattle    GameScreen = 2
)

var (
	gameExitError = errors.New("user exited the game. what a looser")
)

type Game struct {
	screen      GameScreen
	tickCounter uint64

	audioContext *audio.Context
	audioPlayer  *audio.Player

	player   *Player
	entities []*Entity

	droidSelector *Entity
	selectedDroid int

	skipInput bool
}

func NewGame() (g *Game, err error) {
	g = &Game{}

	g.audioContext = audio.NewContext(assets.SampleRate)
	g.audioPlayer, err = g.audioContext.NewPlayer(assets.BackgroundMusic2)
	if err != nil {
		return nil, err
	}
	g.audioPlayer.SetVolume(0.50)
	g.audioPlayer.Play()

	g.player = NewPlayer(g)

	g.droidSelector = NewEntity("droidSelector", assets.UIDroidSelector)
	g.droidSelector.invisible = true
	g.droidSelector.skipInput = true
	g.entities = append(g.entities, g.droidSelector)

	return
}

func (g *Game) Update() error {
	g.tickCounter++
	for _, e := range g.entities {
		e.Tinker()
	}
	switch g.screen {
	case GameScreenTitle:
		for _, e := range g.entities {
			e.Update()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.screen = GameScreenInventory
			g.droidSelector.x = 39
			g.droidSelector.y = 189
			g.showEntity("droidSelector")
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			return gameExitError
		}
	case GameScreenInventory:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.screen = GameScreenBattle
			g.makeAllEntitiesVisible(false)

			d := g.player.inv.droids[g.selectedDroid]
			d.e.x = 286.0
			d.e.y = 668.0
			d.e.invisible = false
			for _, e := range g.entities {
				e.Update()
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.screen = GameScreenTitle
			g.makeAllEntitiesVisible(false)
		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			if x >= 39 && x <= 1260 && y >= 189 && y <= 380 {
				cursor := image.Pt(x, y)
				posx, posy := 39, 189
				for i := 0; i < len(g.player.inv.droids); i++ {
					r := image.Rect(posx, posy, posx+192, posy+192)
					if cursor.In(r) {
						// Select this slot
						g.droidSelector.x = float64(r.Min.X)
						g.droidSelector.y = float64(r.Min.Y)
						g.selectedDroid = i
					}
					posx += 192 + 15 // size + spacing
				}
			}

		} else {
			// Update inventoy item positions to display on inventory screen
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
			for _, e := range g.entities {
				e.Update()
			}
		}
	case GameScreenBattle:
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.screen = GameScreenTitle
			g.makeAllEntitiesVisible(false)
		}
		for _, e := range g.entities {
			e.Update()
		}
	}

	return nil
}

func (g *Game) makeAllEntitiesVisible(visible bool) {
	for _, e := range g.entities {
		e.invisible = !visible
	}
}

func (g *Game) showEntity(name string) {
	for _, e := range g.entities {
		if e.name == name {
			e.invisible = false
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.screen {
	case GameScreenTitle:
		screen.DrawImage(assets.Title, nil)
	case GameScreenInventory:
		screen.DrawImage(assets.InventoryScreen, nil)
	case GameScreenBattle:
		screen.DrawImage(assets.BattleScreen, nil)
	}

	// Draw visible entities
	for _, e := range g.entities {
		e.Draw(screen)
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

	log.Println("Game started")
	if err := ebiten.RunGame(game); err != nil && err != gameExitError {
		log.Fatal(err)
	}
}
