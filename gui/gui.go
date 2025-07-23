package gui

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/op"

	"gioui.org/text"
	"gioui.org/unit"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"gioui.org/x/component"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type C = layout.Context
type D = layout.Dimensions

var th = material.NewTheme()

var text_margins = layout.Inset{
	Top:    unit.Dp(40),
	Bottom: unit.Dp(20),
	Right:  unit.Dp(170),
	Left:   unit.Dp(170),
}

var text_border = widget.Border{
	Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
	CornerRadius: unit.Dp(3),
	Width:        unit.Dp(2),
}

var output_margins = layout.Inset{
	Top:    unit.Dp(0),
	Bottom: unit.Dp(40),
	Right:  unit.Dp(170),
	Left:   unit.Dp(170),
}

var output_border = widget.Border{
	Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
	CornerRadius: unit.Dp(3),
	Width:        unit.Dp(2),
}

var roll_margins = layout.Inset{
	Top:    unit.Dp(0),
	Bottom: unit.Dp(25),
	Right:  unit.Dp(160),
	Left:   unit.Dp(160),
}

var Output_editor = widget.Editor{
	LineHeight: 10,
	ReadOnly:   true,
}

var History_editor = widget.Editor{
	LineHeight: 10,
	ReadOnly:   true,
}

var window_list = widget.List{
	List: layout.List{
		Axis:        layout.Vertical,
		ScrollToEnd: false,
	},
}

// Roll Button
var rollButton widget.Clickable

// Number of Dice
var numDiceInput = widget.Editor{
	SingleLine: true,
	Alignment:  text.Middle,
}

// Type of Dice
var typeDiceInput = widget.Editor{
	SingleLine: true,
	Alignment:  text.Middle,
}

// +/- to Dice Roll
var addRollInput = widget.Editor{
	SingleLine: true,
	Alignment:  text.Middle,
}

// Target Number of Dice
var TargetNumberInput = widget.Editor{
	SingleLine: true,
	Alignment:  text.Middle,
}

// Diffculty of Dice Roll
var diffInput = widget.Editor{
	SingleLine: true,
	Alignment:  text.Middle,
}

func Gui(w *app.Window) error {

	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return appBar.Layout(gtx, th, "string one", "string two")
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Body1(th, "Hello, Gio!").Layout(gtx)
					})
				}),
			)

			materialList := material.List(th, &window_list)

			materialList.Layout(gtx, len(storytellerList), func(gtx C, i int) D {
				return storytellerList[i](gtx)
			})

			e.Frame(gtx.Ops)
		}
	}
}

var modal component.ModalLayer

var appBar = component.AppBar{
	Title:            "go-roll",
	Anchor:           component.Top,
	ModalLayer:       &modal,
	NavigationIcon:   MenuIcon,
	NavigationButton: widget.Clickable{},
}

var MenuIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.NavigationMenu)
	return icon
}()
