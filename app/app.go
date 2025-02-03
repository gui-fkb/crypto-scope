package app

import (
	"crypto-scrope/settings"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	ui *ebitenui.UI

	contentContainer *widget.Container
}

func New() *App {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(settings.BackgroundColor),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Spacing(0, 0),
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch(
				[]bool{true, true, true},
				[]bool{false, true, false}),
		)),
	)

	content := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(settings.BackgroundColor),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	app := &App{
		ui: &ebitenui.UI{
			Container: rootContainer,
		},
		contentContainer: content,
	}

	rootContainer.AddChild(NewStatusBarWidget(), content)

	return app
}

func (g *App) Update() error {
	g.ui.Update()
	return nil
}

func (g *App) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *App) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
