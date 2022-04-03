package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

var SampleRate int = 44100

//go:generate go run ../scripts/generate-assets.go

//go:embed Title.png
var title_png []byte
var Title *ebiten.Image

//go:embed InventoryBackground.png
var inventoryBackground_png []byte
var InventoryBackground *ebiten.Image

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

//go:embed BackgroundMusic.mp3
var backgroundMusic_mp3 []byte
var BackgroundMusic io.Reader

// load loads the image asset as a ebiten.Image pointer.
func load(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func loadMP3(b []byte) (r io.Reader) {
	r, err := mp3.DecodeWithSampleRate(SampleRate, bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	return r
}

func init() {
	Title = load(title_png)
	InventoryBackground = load(inventoryBackground_png)

	BlueL1 = load(blueL1_png)
	BlueL2 = load(blueL2_png)
	BlueL3 = load(blueL3_png)
	BlueL4 = load(blueL4_png)
	BlueL5 = load(blueL5_png)
	BlueL6 = load(blueL6_png)

	BackgroundMusic = loadMP3(backgroundMusic_mp3)
}
