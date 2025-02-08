package main

import (
	"crypto-scrope/app"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

var symbols = []string{
	"btcusdt",
}

const wsBaseUrl = "wss://stream.binance.com:9443"

func main() {
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(getWsSubscriptionStreamUrl(), nil)
	if err != nil {
		panic(err)
	}

	for {
		msgType, bytes, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		fmt.Printf("msgType: %v - msg: %s\n", msgType, string(bytes))

	}

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

func getWsSubscriptionStreamUrl() string {
	var streamNames []string
	for _, symbol := range symbols {
		streamNames = append(streamNames, fmt.Sprintf("%s@depth20", symbol))
	}

	return fmt.Sprintf("%s/stream?streams=%s", wsBaseUrl, strings.Join(streamNames, "/"))
}
