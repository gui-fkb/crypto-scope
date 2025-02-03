package main

import (
	"crypto-scrope/app"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Failed to run the application: %v", err)
	}
}

func run() error {
	if err := configureWindow(); err != nil {
		return err
	}

	application := app.New()

	return ebiten.RunGame(application)
}

func configureWindow() error {
	w, h := ebiten.Monitor().Size()
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Crypto Scope v0.1")
	ebiten.SetWindowPosition(0, 0)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	return nil
}
