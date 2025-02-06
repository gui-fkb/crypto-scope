package app

import (
	"crypto-scrope/settings"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type windowWidget struct {
	*widget.Window
}

func NewWindowWidget(title string) *windowWidget {
	var window windowWidget

	// Construct window content
	windowContentWrapper := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(settings.BackgroundColor),
		),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(),
		),
	)

	windowContent := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(settings.BackgroundColor2),
		),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(
				widget.AnchorLayoutOpts.Padding(widget.Insets{Left: 2, Right: 2, Top: 2}),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: true,
					StretchVertical:   true,
					Padding:           widget.NewInsetsSimple(3),
				},
			),
		),
	)
	windowContentWrapper.AddChild(windowContent)

	// Construct window title
	windowTitleWrapper := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(settings.BackgroundColor),
		),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(0, int(12*settings.Scale)),
		),
	)

	windowTitle := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(settings.BackgroundColor2),
		),
		widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					StretchHorizontal: true,
					StretchVertical:   true,
					Padding:           widget.Insets{Left: 3, Right: 3, Top: 3},
				},
			),
		),
	)

	titleText := widget.NewText(
		widget.TextOpts.Text(title, settings.FontBase, color.White),
		widget.TextOpts.Insets(widget.Insets{Left: int(12 * settings.Scale)}),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionStart,
					VerticalPosition:   widget.AnchorLayoutPositionCenter,
				},
			),
		),
	)

	titleCloseButton := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				Padding:            widget.Insets{Right: int(12 * settings.Scale)},
			}),
		),
		widget.ButtonOpts.Text("x", settings.FontSM, &widget.ButtonTextColor{
			Idle: color.White,
		}),
		widget.ButtonOpts.Image(
			&widget.ButtonImage{
				Idle:    image.NewNineSliceColor(settings.BackgroundColor4),
				Hover:   image.NewNineSliceColor(settings.BackgroundColor5),
				Pressed: image.NewNineSliceColor(settings.BackgroundColor6),
			},
		),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   4,
			Right:  4,
			Top:    1,
			Bottom: 1,
		}),
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				window.Close()
			},
		),
	)

	windowTitle.AddChild(titleText, titleCloseButton)
	windowTitleWrapper.AddChild(windowTitle)

	// Create and return the actual window widget
	window = windowWidget{
		Window: widget.NewWindow(
			widget.WindowOpts.TitleBar(windowTitleWrapper, int(36*settings.Scale)),
			widget.WindowOpts.Contents(windowContentWrapper),
			widget.WindowOpts.Modal(),
			widget.WindowOpts.Draggable(),
			widget.WindowOpts.Resizeable(),
			widget.WindowOpts.MinSize(200, 100),
			widget.WindowOpts.MaxSize(400, 460),
		),
	}

	return &window
}
