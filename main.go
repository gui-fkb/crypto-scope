package main

import (
	"crypto-scrope/app"
	"fmt"
	"log"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/valyala/fastjson"
	"golang.org/x/exp/slices"
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

			var p fastjson.Parser
			v, _ := p.ParseBytes(bytes)

			stream := string(v.GetStringBytes("stream"))

			if strings.Contains(stream, "depth") {
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

			if strings.Contains(stream, "aggTrade") {
				data := v.Get("data")

				eventTime := data.GetInt64("T")
				price, err := strconv.ParseFloat(string(data.GetStringBytes("p")), 64)
				if err != nil {
					slog.Error("error converting price to float")
				}

				time := time.Unix(eventTime/1000, 0)

				quantity, err := strconv.ParseFloat(string(data.GetStringBytes("q")), 64)
				if err != nil {
					slog.Error("error converting quantity to float")
				}

				if len(app.MarketTrades) < 15 {
					app.MarketTrades = append(app.MarketTrades, app.MarketTrade{
						Price:    price,
						Quantity: quantity,
						Time:     time,
					})
				} else {
					slices.Delete(app.MarketTrades, 0, 1)
					slices.Insert(app.MarketTrades, 14, app.MarketTrade{
						Price:    price,
						Quantity: quantity,
						Time:     time,
					})
				}

				fmt.Printf("eventTime: %s, price: %f, quantity: %f\n", time.String(), price, quantity)
				fmt.Printf("market trades: %v\n", len(app.MarketTrades))
			}
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
		streamNames = append(streamNames, fmt.Sprintf("%s@aggTrade", symbol))
	}

	return fmt.Sprintf("%s/stream?streams=%s", wsBaseUrl, strings.Join(streamNames, "/"))
}
