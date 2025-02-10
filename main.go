package main

import (
	"crypto-scrope/app"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/valyala/fastjson"
)

var symbols = []string{
	"btcusdt",
}

const wsBaseUrl = "wss://stream.binance.com:9443"

func main() {
	go func() {
		dialer := websocket.DefaultDialer
		conn, _, err := dialer.Dial(getWsSubscriptionStreamUrl(), nil)
		if err != nil {
			panic(err)
		}

		for {
			_, bytes, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}

			//fmt.Printf("msgType: %v - msg: %s\n", msgType, string(bytes))

			var p fastjson.Parser

			v, err := p.ParseBytes(bytes)

			data := v.Get("data")

			bids := data.GetArray("bids")
			asks := data.GetArray("asks")

			var bidSlice []app.OrderBookData
			var askSlice []app.OrderBookData

			for _, v := range bids {
				price, err := strconv.ParseFloat(string(v.GetStringBytes("0")), 10)
				if err != nil {
					fmt.Println("error converting price to float")
					break
				}
				quantity, err := strconv.ParseFloat(string(v.GetStringBytes("1")), 10)
				if err != nil {
					fmt.Println("error converting quantity to float")
					break
				}

				obData := app.OrderBookData{
					Price:    price,
					Quantity: quantity,
					Sum:      price * quantity,
				}

				bidSlice = append(bidSlice, obData)
			}

			for _, v := range asks {
				price, err := strconv.ParseFloat(string(v.GetStringBytes("0")), 10)
				if err != nil {
					fmt.Println("error converting price to float")
					break
				}
				quantity, err := strconv.ParseFloat(string(v.GetStringBytes("1")), 10)
				if err != nil {
					fmt.Println("error converting quantity to float")
					break
				}

				obData := app.OrderBookData{
					Price:    price,
					Quantity: quantity,
					Sum:      price * quantity,
				}

				askSlice = append(askSlice, obData)
			}

			app.Ob.Bids = bidSlice
			app.Ob.Asks = askSlice
		}
	}()

	err := run()
	if err != nil {
		log.Fatalf("could not run the application: %v", err)
	}
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
		streamNames = append(streamNames, fmt.Sprintf("%s@depth20@100ms", symbol))
	}

	return fmt.Sprintf("%s/stream?streams=%s", wsBaseUrl, strings.Join(streamNames, "/"))
}
