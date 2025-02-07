package app

import (
	"crypto-scrope/settings"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type orderBookWidget struct {
	*widget.Container

	rows []*orderBookRowWidget
}

type orderBookRowWidget struct {
	*widget.Container
}

func NewOrderBookWidget() *orderBookWidget {
	rowCount := 15

	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(settings.BackgroundColor),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Spacing(4, 4),
			widget.GridLayoutOpts.Stretch([]bool{true}, benerateBoolArray(rowCount, true)),
		)),

		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchVertical: true,
				},
			),
		),
	)

	var rows []*orderBookRowWidget

	for i := 0; i < rowCount; i++ {
		obRow := NewOrderBookRowWidget()
		rows = append(rows, obRow)

		container.AddChild(obRow)
	}

	return &orderBookWidget{
		Container: container,
		rows:      rows,
	}
}

func NewOrderBookRowWidget() *orderBookRowWidget {
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(settings.BackgroundColor)),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(3),
			widget.GridLayoutOpts.Spacing(int(4*settings.Scale), int(10*settings.Scale)),
			widget.GridLayoutOpts.Stretch([]bool{true, true, true}, []bool{true}),
		)),
	)

	contentA := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(settings.Red)),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: true,
					StretchVertical:   true,
				},
			),
			widget.WidgetOpts.MinSize(0, int(22*settings.Scale)),
		),
	)

	contentB := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(settings.Red)),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: true,
					StretchVertical:   true,
				},
			),
			widget.WidgetOpts.MinSize(0, int(22*settings.Scale)),
		),
	)

	contentC := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(settings.Red)),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: true,
					StretchVertical:   true,
				},
			),
			widget.WidgetOpts.MinSize(0, int(22*settings.Scale)),
		),
	)

	container.AddChild(contentA, contentB, contentC)

	return &orderBookRowWidget{
		Container: container,
	}
}

func benerateBoolArray(length int, value bool) []bool {
	result := make([]bool, length)
	for i := range result {
		result[i] = value
	}
	return result
}
