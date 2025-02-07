package main

import (
	"crypto-scrope/actor"
	"crypto-scrope/app"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	actor.RunActor()

	// err := run()
	// if err != nil {
	// 	log.Fatalf("could not run the application: %v", err)
	// }
}

func run() error {
	setupWindow()

	application := app.New()

	return ebiten.RunGame(application)
}

func setupWindow() {
	w, h := ebiten.Monitor().Size()
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("crypto scope v0.0.1")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
}
