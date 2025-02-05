package app

import (
	"crypto-scrope/settings"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type menuBarWidget struct {
	*widget.Container
}

func NewMenuBarWidget() *menuBarWidget {
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(settings.BackgroundColor2)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(0, int(40*settings.Scale)),
		),
	)

	return &menuBarWidget{
		Container: container,
	}
}

func (w *menuBarWidget) PreferredSize() (int, int) {
	return 0, int(40 * settings.Scale)
}
