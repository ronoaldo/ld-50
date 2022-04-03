package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:generate go run ../scripts/generate-assets.go

//go:embed Title.png
var title_png []byte
var Title *ebiten.Image

//go:embed BlueL1.png
var blueL1_png []byte
var BlueL1 *ebiten.Image

// load loads the image asset as a ebiten.Image pointer.
func load(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func init() {
	BlueL1 = load(blueL1_png)
	Title = load(title_png)
}
