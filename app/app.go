package app

import (
	"crypto-scrope/settings"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type App struct {
	ui *ebitenui.UI
}

func New() *App {
	// construct a new container that will serve as the root of the UI hierarchy
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

	contentContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(settings.BackgroundColor),
		),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(),
		),
	)

	rootContainer.AddChild(NewMenuBarWidget(), contentContainer, NewFooterWidget())

	application := App{
		ui: &ebitenui.UI{
			Container: rootContainer,
		},
	}

	return &application
}

// Draw implements ebiten.Game.
func (g *App) Draw(screen *ebiten.Image) {
	// Clear the screen with the color teal
	screen.Fill(colornames.Teal)

	// Draw the UI onto the screen
	g.ui.Draw(screen)
}

// Layout implements ebiten.Game.
func (g *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Update implements ebiten.Game.
func (g *App) Update() error {
	// Update the UI
	g.ui.Update()

	return nil
}
