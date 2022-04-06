package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
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
	GameOverScreen      GameScreen = 3
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

	inputSelector *Entity
	selectedDroid int
	selectedSkill int

	playerUnit   *BattleUnit
	playerTimer  time.Time
	isPlayerTurn bool
	enemyUnit    *BattleUnit
	enemyEntity  *Entity

	gameOver bool
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

	g.inputSelector = NewEntity("inputSelector", assets.UIDroidSelector)
	g.inputSelector.invisible = true
	g.inputSelector.skipInput = true
	g.entities = append(g.entities, g.inputSelector)

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
			g.inputSelector.x = 39
			g.inputSelector.y = 189
			g.showEntity("inputSelector")
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			return gameExitError
		}
	case GameScreenInventory:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.screen = GameScreenBattle
			g.makeAllEntitiesVisible(false)
			g.showEntity("inputSelector")
			g.inputSelector.x = 657
			g.inputSelector.y = 817

			d := g.player.inv.droids[g.selectedDroid]
			d.e.x = 226.0
			d.e.y = 510.0
			d.e.invisible = false
			d.e.skipInput = true

			g.playerUnit = &BattleUnit{
				Name:   d.Name,
				Stats:  d.Stats(),
				Skills: d.Skills,
			}
			g.playerTimer = time.Now()
			g.isPlayerTurn = true

			g.enemyEntity = NewEntity("Enemy", assets.OctopusEnemi)
			g.enemyEntity.x, g.enemyEntity.y = 1505.0, 513.0
			g.enemyEntity.skipInput = true
			g.entities = append(g.entities, g.enemyEntity)

			g.enemyUnit = &BattleUnit{
				Name:   "Clone Droid",
				Stats:  d.Stats(),
				Skills: d.Skills,
			}

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
						g.inputSelector.x = float64(r.Min.X)
						g.inputSelector.y = float64(r.Min.Y)
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
		if g.playerUnit.Stats.HP <= 0 || g.enemyUnit.Stats.HP <= 0 {
			g.screen = GameOverScreen
			g.makeAllEntitiesVisible(false)
			return nil
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.screen = GameScreenTitle
			g.makeAllEntitiesVisible(false)
		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			if x >= 657 && x <= 1440 && y >= 800 && y <= 1100 {
				cursor := image.Pt(x, y)
				posx, posy := 657, 817
				for i := 0; i < 3; i++ {
					r := image.Rect(posx, posy, posx+192, posy+192)
					if cursor.In(r) {
						// Select this slot
						g.inputSelector.x = float64(r.Min.X)
						g.inputSelector.y = float64(r.Min.Y)
						g.selectedSkill = i
					}
					posx += 192 + 15 // size + spacing
				}
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			if g.selectedSkill > 0 {
				// Select this slot
				g.inputSelector.x -= (192 + 15)
				g.selectedSkill -= 1
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			if g.selectedSkill < 2 {
				// Select this slot
				g.inputSelector.x += (192 + 15)
				g.selectedSkill += 1
			}
		}

		// Timer logic
		if time.Since(g.playerTimer) > 10*time.Second || inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			// Time expired for player to choose the skill
			skill := g.playerUnit.Skills[g.selectedSkill]
			skill.Effect(g.playerUnit, g.enemyUnit)
			log.Printf("Player action: %s", skill.Name)
			g.playerTimer = time.Now()
			g.isPlayerTurn = false
		} else if !g.isPlayerTurn {
			// Other player logic
			// AI of the enemy - choose a random skill
			skill := g.playerUnit.Skills[rand.Int()%3]
			skill.Effect(g.enemyUnit, g.playerUnit)
			log.Printf("Enemy action: %s", skill.Name)
			g.playerTimer = time.Now()
			g.isPlayerTurn = true
		}

		for _, e := range g.entities {
			e.Update()
		}
	case GameOverScreen:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.screen = GameScreenTitle
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

		// Draw Player HP
		hpText := fmt.Sprintf("HP: %d\n%s", g.playerUnit.Stats.HP, g.player.game.playerUnit.Name)
		text.Draw(screen, hpText, assets.RobotoMonoRegular, 132, 64, color.White)

		// Draw Enemy HP
		hpText = fmt.Sprintf("HP: %d\n%s", g.enemyUnit.Stats.HP, g.enemyEntity.name)
		text.Draw(screen, hpText, assets.RobotoMonoRegular, 1406, 64, color.White)

		// Draw turn timer
		if !g.gameOver {
			timerText := fmt.Sprintf("%v", time.Since(g.playerTimer))
			text.Draw(screen, timerText, assets.RobotoMonoRegular, 798, 130, color.White)
		}
	case GameOverScreen:
		screen.DrawImage(assets.GameOverScreen, nil)
		result := "You Lost!"
		if g.playerUnit.Stats.HP > 0 {
			result = "You won!"
		}
		text.Draw(screen, result, assets.RobotoMonoRegular, 863, 718, color.White)
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
