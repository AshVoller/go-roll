package gui

import (
	"image/color"
	"log"

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

// TODO add logger

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

// App Bar and Nav Menu
var modal = component.NewModal()
var appBar = component.NewAppBar(modal)
var modalSideDraw = component.NewModalNav(modal, "Dice Systems", "Choose which system to use.")

// App Bar Event Handler
func appBarEvents(gtx layout.Context) {
	for _, navi_event := range appBar.Events(gtx) {
		switch navi_event.(type) {
		case component.AppBarNavigationClicked:
			modalSideDraw.Appear(gtx.Now)
			// log.Printf("button pushed: %v", n)
		case component.AppBarContextMenuDismissed:
			// log.Printf("Context Menu Dismissed: %v", n)
		case component.AppBarOverflowActionClicked:
			// log.Printf("Overflow Action Clicked: %v", n)
		}
	}
}

// Navi Menu Option to Dice Roller Content List
type listID int

const (
	basicConst listID = iota
	storytellerConst
)

func GetList(id listID) []layout.Widget {
	switch id {
	case basicConst:
		// log.Printf("the list is: %v", basicList)
		return basicList
	case storytellerConst:
		// log.Printf("the list is: %v", storytellerList)
		return storytellerList
	default:
		return nil
	}
}

// Dice Roller Content List
var ContentList []layout.Widget = basicList

func content(gtx C, th *material.Theme, contentList []layout.Widget) D {
	materialList := material.List(th, &window_list)

	return materialList.Layout(gtx, len(contentList), func(gtx C, i int) D {
		return contentList[i](gtx)
	})
}

func Gui(w *app.Window) error {

	var ops op.Ops

	// appBar
	var MenuIcon *widget.Icon = func() *widget.Icon {
		icon, _ := widget.NewIcon(icons.NavigationMenu)
		return icon
	}()
	appBar.NavigationIcon = MenuIcon
	appBar.Title = "go-roll"
	appBar.ContextualTitle = "Contextual Menu"
	appBar.Anchor = component.Top

	// Nav Menu Options
	modalSideDraw.AddNavItem(
		component.NavItem{
			Tag:  0,
			Name: "Basic Dice",
			// Icon: MenuIcon,
		},
	)

	modalSideDraw.AddNavItem(
		component.NavItem{
			Tag:  1,
			Name: "Exalted",
			// Icon: MenuIcon,
		},
	)

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			appBarEvents(gtx)

			if modalSideDraw.NavDestinationChanged() {
				currentNav := modalSideDraw.CurrentNavDestination()
				tag, err := currentNav.(int)
				if !err {
					log.Printf("Current Nav is not a int: %v", err)
				}

				var id listID = listID(tag)
				ContentList = GetList(id)
			}

			bar := layout.Rigid(func(gtx C) D {
				// TODO figure out "string A" and "string B" acessibility nav option
				return appBar.Layout(gtx, th, "string A", "string B")
			})

			menu := layout.Flexed(1, func(gtx C) D {
				return content(gtx, th, ContentList)
			})

			layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
			}.Layout(gtx, bar, menu)

			modal.Layout(gtx, th)

			e.Frame(gtx.Ops)
		}
	}
}
