package settings

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	// Ebiten
	Scale = float32(ebiten.Monitor().DeviceScaleFactor())

	// Colors
	Black          = color.RGBA{12, 14, 17, 255}
	Red            = color.RGBA{246, 71, 93, 255}
	Green          = color.RGBA{45, 189, 133, 255}
	OrderbookRed   = color.RGBA{53, 20, 29, 255}
	OrderbookGreen = color.RGBA{27, 45, 43, 1}

	BackgroundColor  = Black
	BackgroundColor2 = color.RGBA{23, 26, 32, 255}
	BackgroundColor3 = color.RGBA{23, 26, 32, 0}

	BackgroundColor4 = color.RGBA{30, 33, 24, 1}
	BackgroundColor5 = color.RGBA{45, 50, 35, 1}
	BackgroundColor6 = color.RGBA{20, 22, 16, 1}

	FontSM   text.Face
	FontBase text.Face
)

func init() {
	FontSM, _ = LoadFont(12)
	FontBase, _ = LoadFont(13)
}

func LoadFont(size float64) (text.Face, error) {
	b, err := os.Open("assets/jetbrains.ttf")
	if err != nil {
		return nil, err
	}
	s, err := text.NewGoTextFaceSource(b)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &text.GoTextFace{
		Source: s,
		Size:   size * ebiten.Monitor().DeviceScaleFactor(),
	}, nil
}

func ColorWithAlpha(c color.RGBA, a uint8) color.RGBA {
	c.A = a
	return c
}
