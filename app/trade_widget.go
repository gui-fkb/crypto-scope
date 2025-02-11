package app

import (
	"crypto-scrope/app/helper"
	"crypto-scrope/settings"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type tradeWidget struct {
	*widget.Container

	rows []*tradeRowWidget
}

type tradeRowWidget struct {
	*widget.Container

	price    *widget.Text
	quantity *widget.Text
	date     *widget.Text
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

	return &tradeWidget{
		Container: container,
		rows:      rows,
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
		date:     textC,
	}
}
