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

//go:embed BlueL1.png
var blueL1_png []byte
var BlueL1 *ebiten.Image

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
	BlueL1 = load(blueL1_png)
	Title = load(title_png)
	BackgroundMusic = loadMP3(backgroundMusic_mp3)
}
