package main

import (
	"crypto-scrope/app"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	w, h := ebiten.Monitor().Size()
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("crypto scope v0.0.1")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	application := app.New()

	err := ebiten.RunGame(application)
	if err != nil {
		log.Fatal(err)
	}
}
