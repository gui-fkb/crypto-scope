package app

import (
	"crypto-scrope/settings"
	"fmt"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type StatusBarWidget struct {
	*widget.Container

	fpsLabel *widget.Text
}

func NewStatusBarWidget() *StatusBarWidget {
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(settings.PanelBackgroundColor),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.Insets{Left: int(settings.PanelPadding)}),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(0, int(settings.AppFooterHeight)),
		),
	)
	fpsLabel := widget.NewText(
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				HorizontalPosition: widget.AnchorLayoutPositionStart,
			}),
		),
		widget.TextOpts.Text("60", settings.FontSM, color.White),
	)
	container.AddChild(fpsLabel)

	return &StatusBarWidget{
		Container: container,
		fpsLabel:  fpsLabel,
	}
}

func (w *StatusBarWidget) Render(screen *ebiten.Image) {
	w.Container.Render(screen)

	fps := ebiten.ActualFPS()
	w.fpsLabel.Label = fmt.Sprintf("FPS %d", int(fps))
}

func (w *StatusBarWidget) PreferredSize() (int, int) {
	return 0, int(settings.AppFooterHeight)
}
