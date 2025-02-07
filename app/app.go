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
	ui      *ebitenui.UI
	content *widget.Container
}

func New() *App {
	ui := &ebitenui.UI{}

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

	ui.Container = rootContainer
	application := App{
		ui:      ui,
		content: contentContainer,
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

var (
	timer   float64
	elapsed bool
)

// Update implements ebiten.Game.
func (app *App) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	app.ui.Update()

	if !elapsed {
		timer += 1.0 / 60.0
		if timer > 1 {
			app.createDefaultLayout()
			elapsed = true
		}
	}

	return nil
}

func (app *App) createDefaultLayout() {
	window := NewWindowWidget("binancef - btcustd", NewOrderBookWidget()).Window

	r := app.content.GetWidget().Rect
	window.SetLocation(r)
	app.ui.AddWindow(window)
}
