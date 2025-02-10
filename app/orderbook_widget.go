package app

import (
	"crypto-scrope/settings"
	"fmt"
	"image/color"
	"slices"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var Ob OrderBook

type OrderBook struct {
	Bids []OrderBookData
	Asks []OrderBookData
}

type OrderBookData struct {
	Price    float64
	Quantity float64
	Sum      float64
}

type orderBookWidget struct {
	*widget.Container

	rows []*orderBookRowWidget
}

type orderBookRowWidget struct {
	*widget.Container

	price    *widget.Text
	quantity *widget.Text
	sum      *widget.Text
}

func NewOrderBookWidget() *orderBookWidget {
	rowCount := 14

	container := widget.NewContainer(
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
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.Transparent)),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(3),
			widget.GridLayoutOpts.Spacing(int(24*settings.Scale), int(10*settings.Scale)),
			widget.GridLayoutOpts.Stretch([]bool{true, true, true}, []bool{true}),
		)),
	)

	contentA := widget.NewContainer(
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

	textA := widget.NewText(
		widget.TextOpts.Text("....", settings.FontSM, color.White),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
	)
	contentA.AddChild(textA)

	contentB := widget.NewContainer(
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

	textB := widget.NewText(
		widget.TextOpts.Text("44.440", settings.FontSM, color.White),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionStart),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
	)
	contentB.AddChild(textB)

	contentC := widget.NewContainer(
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

	textC := widget.NewText(
		widget.TextOpts.Text("44.440", settings.FontSM, color.White),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionStart),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
	)
	contentC.AddChild(textC)

	container.AddChild(contentA, contentB, contentC)

	return &orderBookRowWidget{
		Container: container,
		price:     textA,
		quantity:  textB,
		sum:       textC,
	}
}

func (w *orderBookWidget) Render(screen *ebiten.Image) {
	if len(Ob.Bids) == 0 || len(Ob.Asks) == 0 {
		return
	}

	maxAskSum := slices.MaxFunc(Ob.Asks, func(a, b OrderBookData) int {
		if a.Sum > b.Sum {
			return 1
		}
		return -1
	}).Sum

	for i := range Ob.Asks {
		if i > 6 {
			break
		}

		sum := Ob.Asks[6-i].Sum
		w.rows[i].price.Label = fmt.Sprintf("%.2f", Ob.Asks[6-i].Price)
		w.rows[i].quantity.Label = fmt.Sprintf("%.5f", Ob.Asks[6-i].Quantity)
		w.rows[i].sum.Label = formatWithK(sum)

		w.rows[i].price.Color = settings.Red
		w.rows[i].quantity.Color = settings.Red
		w.rows[i].sum.Color = settings.Red

		bgContainer := w.rows[i].Container
		fillPerc := sum / maxAskSum

		vector.DrawFilledRect(screen, float32(bgContainer.GetWidget().Rect.Min.X), float32(bgContainer.GetWidget().Rect.Min.Y), float32(bgContainer.GetWidget().Rect.Dx())*float32(fillPerc), float32(bgContainer.GetWidget().Rect.Dy()), settings.OrderbookRed, false)
	}

	maxBidSum := slices.MaxFunc(Ob.Bids, func(a, b OrderBookData) int {
		if a.Sum > b.Sum {
			return 1
		}
		return -1
	}).Sum

	for i, bid := range Ob.Bids {
		if i > 6 {
			break
		}

		sum := bid.Sum
		w.rows[i+7].price.Label = fmt.Sprintf("%.2f", bid.Price)
		w.rows[i+7].quantity.Label = fmt.Sprintf("%.5f", bid.Quantity)
		w.rows[i+7].sum.Label = formatWithK(sum)

		w.rows[i+7].price.Color = settings.Green
		w.rows[i+7].quantity.Color = settings.Green
		w.rows[i+7].sum.Color = settings.Green

		bgContainer := w.rows[i+7].Container
		fillPerc := sum / maxBidSum

		vector.DrawFilledRect(screen, float32(bgContainer.GetWidget().Rect.Min.X), float32(bgContainer.GetWidget().Rect.Min.Y), float32(bgContainer.GetWidget().Rect.Dx())*float32(fillPerc), float32(bgContainer.GetWidget().Rect.Dy()), settings.OrderbookGreen, false)
	}

	w.Container.Render(screen)
}

func benerateBoolArray(length int, value bool) []bool {
	result := make([]bool, length)
	for i := range result {
		result[i] = value
	}
	return result
}

func formatWithK(value float64) string {
	if value >= 1000 {
		return fmt.Sprintf("%.2fK", value/1000)
	}
	return fmt.Sprintf("%.5f", value)
}
