package assets

import (
	"bytes"
	"image"
	"io"
	"io/ioutil"
	"log"
	"time"

	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var SampleRate int = 44100

//go:generate go run ../cmd/genassets/main.go

//go:embed TitleScreen.png
var title_png []byte
var Title *ebiten.Image

//go:embed InventoryScreen.png
var inventoryScreen_png []byte
var InventoryScreen *ebiten.Image

//go:embed BattleScreen.png
var battleScreen_png []byte
var BattleScreen *ebiten.Image

//go:embed UIDroidSelector.png
var uIDroidSelector_png []byte
var UIDroidSelector *ebiten.Image

//go:embed fonts/RobotoMono-Regular.ttf
var rotoboMonoRegular_ttf []byte
var RobotoMonoRegular font.Face

//go:embed BlueL1.png
var blueL1_png []byte
var BlueL1 *ebiten.Image

//go:embed BlueL2.png
var blueL2_png []byte
var BlueL2 *ebiten.Image

//go:embed BlueL3.png
var blueL3_png []byte
var BlueL3 *ebiten.Image

//go:embed BlueL4.png
var blueL4_png []byte
var BlueL4 *ebiten.Image

//go:embed BlueL5.png
var blueL5_png []byte
var BlueL5 *ebiten.Image

//go:embed BlueL6.png
var blueL6_png []byte
var BlueL6 *ebiten.Image

//go:embed chip-life.png
var chipLife_png []byte
var ChipLife *ebiten.Image

//go:embed chip-speed.png
var chipSpeed_png []byte
var ChipSpeed *ebiten.Image

//go:embed chip-strength.png
var chipStrength_png []byte
var ChipStrength *ebiten.Image

//go:embed BackgroundMusic.mp3
var backgroundMusic_mp3 []byte
var BackgroundMusic io.ReadSeeker

//go:embed BackgroundMusic2.mp3
var backgroundMusic2_mp3 []byte
var BackgroundMusic2 io.ReadSeeker

//go:embed octopus-enemi.png
var octopusEnemi_png []byte
var OctopusEnemi *ebiten.Image

// load loads the image asset as a ebiten.Image pointer.
func load(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func loadMP3(b []byte, infinite bool) io.ReadSeeker {
	start := time.Now()
	m, err := mp3.DecodeWithSampleRate(SampleRate, bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	if infinite {
		// We need to calculate the length after decoding.
		d, _ := ioutil.ReadAll(m)
		// Reset
		m.Seek(0, io.SeekStart)
		return audio.NewInfiniteLoop(m, int64(len(d)))
	}
	log.Printf("Loaded MP3 (%d bytes) took %v", len(b), time.Since(start))
	return m
}

func loadFont(font []byte) font.Face {
	tt, err := opentype.Parse(font)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 32,
		DPI:  72,
	})

	if err != nil {
		panic(err)
	}

	return face
}

func init() {
	Title = load(title_png)
	InventoryScreen = load(inventoryScreen_png)
	BattleScreen = load(battleScreen_png)
	UIDroidSelector = load(uIDroidSelector_png)

	RobotoMonoRegular = loadFont(rotoboMonoRegular_ttf)

	BlueL1 = load(blueL1_png)
	BlueL2 = load(blueL2_png)
	BlueL3 = load(blueL3_png)
	BlueL4 = load(blueL4_png)
	BlueL5 = load(blueL5_png)
	BlueL6 = load(blueL6_png)

	ChipLife = load(chipLife_png)
	ChipStrength = load(chipStrength_png)
	ChipSpeed = load(chipSpeed_png)

	OctopusEnemi = load(octopusEnemi_png)

	BackgroundMusic = loadMP3(backgroundMusic_mp3, true)
	BackgroundMusic2 = loadMP3(backgroundMusic2_mp3, true)
}
