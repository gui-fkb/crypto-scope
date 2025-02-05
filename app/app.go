package app

import (
	"crypto-scrope/settings"
	"image/color"

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
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(settings.BackgroundColor)),

		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(30)),
			widget.GridLayoutOpts.Spacing(20, 10),
			widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{false, true}),
		)),
	)

	// now let's contruct the inner containers
	innerContainer1 := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(settings.BackgroundColor2)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(50, 50),
		),
	)

	innerContainer2 := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(settings.Red)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(50, 50),
		),
	)

	innerContainer3 := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(settings.Green)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{
				HorizontalPosition: widget.GridLayoutPositionCenter,
				VerticalPosition:   widget.GridLayoutPositionCenter,
				MaxWidth:           100,
				MaxHeight:          100,
			}),
			widget.WidgetOpts.MinSize(50, 50),
		),
	)

	innerContainer4 := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 255, 255, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(50, 50),
		),
	)

	rootContainer.AddChild(
		innerContainer1,
		innerContainer2,
		innerContainer3,
		innerContainer4,
	)

	ui := ebitenui.UI{
		Container: rootContainer,
	}

	application := App{
		ui: &ui,
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
