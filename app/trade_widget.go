package app

import (
	"crypto-scrope/app/helper"
	"crypto-scrope/settings"
	"fmt"
	"image/color"
	"time"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

var MarketTrades []MarketTrade

type MarketTrade struct {
	Price    float64
	Quantity float64
	Time     time.Time
}

type tradeWidget struct {
	*widget.Container

	rows  []*tradeRowWidget
	flash *ebiten.Image
}

type tradeRowWidget struct {
	*widget.Container

	price    *widget.Text
	quantity *widget.Text
	time     *widget.Text
}

func NewTradeWidget() *tradeWidget {
	rowNumber := 15

	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.Transparent)),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Spacing(4, 4),
				widget.GridLayoutOpts.Stretch([]bool{true, true, true}, helper.GenerateBoolArray(rowNumber, true)),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchVertical: true,
				},
			),
		),
	)

	var rows []*tradeRowWidget
	for i := 0; i < rowNumber; i++ {
		tRow := NewTradeRowWidget()
		rows = append(rows, tRow)

		container.AddChild(tRow)
	}

	flash := ebiten.NewImage(1, 1)
	flash.Fill(color.RGBA{255, 255, 255, 128})

	return &tradeWidget{
		Container: container,
		rows:      rows,
		flash:     flash,
	}
}

func NewTradeRowWidget() *tradeRowWidget {
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.Transparent)),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(3),
				widget.GridLayoutOpts.Spacing(int(24*settings.Scale), int(10*settings.Scale)),
				widget.GridLayoutOpts.Stretch([]bool{true, true, true}, []bool{true}),
			),
		),
	)

	containerA := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.Transparent)),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	containerB := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.Transparent)),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	containerC := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.Transparent)),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	textA := widget.NewText(
		widget.TextOpts.Text(":....", settings.FontSM, color.White),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
	)

	textB := widget.NewText(
		widget.TextOpts.Text("44.440", settings.FontSM, color.White),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
	)

	textC := widget.NewText(
		widget.TextOpts.Text("44.440", settings.FontSM, color.White),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
	)

	containerA.AddChild(textA)
	containerB.AddChild(textB)
	containerC.AddChild(textC)

	container.AddChild(containerA, containerB, containerC)

	return &tradeRowWidget{
		Container: container,

		price:    textA,
		quantity: textB,
		time:     textC,
	}
}

var lastPrice float64

func (w *tradeWidget) Render(screen *ebiten.Image) {
	if len(MarketTrades) < 15 {
		w.Container.Render(screen)
		return
	}

	for i, mt := range MarketTrades {
		w.rows[14-i].price.Label = fmt.Sprintf("%.2f", mt.Price)
		w.rows[14-i].quantity.Label = fmt.Sprintf("%.5f", mt.Quantity)
		w.rows[14-i].time.Label = mt.Time.Format(time.TimeOnly)

		if i > 0 && mt.Price >= MarketTrades[i-1].Price {
			w.rows[14-i].price.Color = settings.Green
			w.rows[14-i].quantity.Color = settings.Green
		} else {
			w.rows[14-i].price.Color = settings.Red
			w.rows[14-i].quantity.Color = settings.Red
		}
	}

	if lastPrice != MarketTrades[14].Price {
		rect := w.rows[0].GetWidget().Rect

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(rect.Dx()), float64(rect.Dy()))
		op.GeoM.Translate(float64(rect.Min.X), float64(rect.Min.Y))
		op.Blend = ebiten.BlendSourceOver
		screen.DrawImage(w.flash, op)
	}

	lastPrice = MarketTrades[14].Price
	w.Container.Render(screen)
}
