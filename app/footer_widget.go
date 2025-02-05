package app

import (
	"crypto-scrope/settings"
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type footerWidget struct {
	*widget.Container

	fpsLabel *widget.Text
}

func NewFooterWidget() *footerWidget {
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(settings.BackgroundColor2),
		),

		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.Insets{Left: int(12 * settings.Scale)}),
		)),

		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(0, int(40*settings.Scale)),
		),
	)

	fpsLabel := widget.NewText(
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.TextOpts.Text("60", settings.FontSM, color.White),
	)

	container.AddChild(fpsLabel)

	return &footerWidget{
		Container: container,
		fpsLabel:  fpsLabel,
	}
}

func (w *footerWidget) Render(screen *ebiten.Image) {
	fps := ebiten.ActualFPS()
	w.fpsLabel.Label = fmt.Sprintf("FPS %d", int(fps))

	w.Container.Render(screen)
}

func (w *footerWidget) PreferredSize() (int, int) {
	return 0, int(40 * settings.Scale)
}
