package gui

import (
	"image/color"
	"log"
	"time"

	"gioui.org/app"
	"gioui.org/op"
	"golang.org/x/exp/shiny/materialdesign/icons"

	"gioui.org/text"
	"gioui.org/unit"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"gioui.org/x/component"
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

	appBar.NavigationIcon = MenuIcon
	appBar.Title = "go-roll"
	appBar.ContextualTitle = "Contextual Menu"
	appBar.Anchor = component.Top

	modalSideDraw.AddNavItem(
		component.NavItem{
			Tag:  "3",
			Name: "Test 3",
			Icon: MenuIcon,
		},
	)

	modalSideDraw.AddNavItem(
		component.NavItem{
			Tag:  "4",
			Name: "Test 4",
			Icon: MenuIcon,
		},
	)

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			appBarEvents(gtx)

			bar := layout.Rigid(func(gtx C) D {
				return appBar.Layout(gtx, th, "string A", "string B")
			})

			menu := layout.Flexed(1, func(gtx C) D {
				return content(gtx, th)
			})

			flex := layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
			}
			flex.Layout(gtx, bar, menu)

			modal.Layout(gtx, th)

			e.Frame(gtx.Ops)
		}
	}
}

func appBarEvents(gtx layout.Context) {
	for _, navi_event := range appBar.Events(gtx) {
		switch n := navi_event.(type) {
		case component.AppBarNavigationClicked:
			modalSideDraw.Appear(gtx.Now)
			sideAnim.Disappear(gtx.Now)
			log.Printf("button pushed: %v", n)
		case component.AppBarContextMenuDismissed:
			log.Printf("Context Menu Dismissed: %v", n)
		case component.AppBarOverflowActionClicked:
			log.Printf("Overflow Action Clicked: %v", n)
		}
	}
}

func content(gtx C, th *material.Theme) D {
	materialList := material.List(th, &window_list)

	return materialList.Layout(gtx, len(storytellerList), func(gtx C, i int) D {
		return storytellerList[i](gtx)
	})
}

var sideAnim = component.VisibilityAnimation{
	State:    component.Invisible,
	Duration: time.Millisecond * 250,
}

var modal = component.NewModal()

var modalSideDraw = component.NewModalNav(modal, "Option 2", "Subtitle 2")

var appBar = component.NewAppBar(modal)

var MenuIcon *widget.Icon = func() *widget.Icon {
	icon, _ := widget.NewIcon(icons.NavigationMenu)
	return icon
}()
