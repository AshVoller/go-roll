package gui

import (
	"go-roll/roller"
	"image/color"

	"gioui.org/app"
	"gioui.org/op"

	"gioui.org/text"
	"gioui.org/unit"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

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

func Gui(w *app.Window) error {
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

	var ops op.Ops

	th := material.NewTheme()

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			listItems := []layout.Widget{
				func(gtx C) D {
					// Flex Box of the whole page
					return layout.Flex{
						Axis:      layout.Vertical,
						Alignment: layout.Middle,
						//Spacing:   layout.SpaceStart, // come back to later
					}.Layout(gtx,
						layout.Rigid(layout.Spacer{Height: unit.Dp(50)}.Layout),

						layout.Rigid(func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Dp(100)
							gtx.Constraints.Max.X = gtx.Dp(300)
							return material.Body2(th, "Roll Output").Layout(gtx)
						}),

						layout.Rigid(func(gtx C) D {
							return output_margins.Layout(gtx, func(gtx C) D {
								return output_border.Layout(gtx, func(gtx C) D {
									gtx.Constraints.Min.Y = gtx.Dp(250)
									gtx.Constraints.Max.Y = gtx.Dp(250)
									return material.Editor(th, &Output_editor, "Enter a number of dice,\nNumber of sides of dice,\nand +/- to the dice roll.").Layout(gtx)
								})
							})
						}),

						// Rigid Box of the Horizontal Text Inputs
						layout.Rigid(func(gtx C) D {
							return layout.Center.Layout(gtx, func(gtx C) D {
								return text_margins.Layout(gtx, func(gtx C) D {
									return layout.Flex{
										Axis: layout.Horizontal,
										// Alignment: layout.Middle,
									}.Layout(gtx,
										layout.Rigid(func(gtx C) D {
											gtx.Constraints.Min.X = gtx.Dp(100)
											gtx.Constraints.Max.X = gtx.Dp(300)
											return material.Body2(th, "Number of Dice").Layout(gtx)
										}),

										layout.Rigid(layout.Spacer{Width: 5}.Layout),
										layout.Rigid(func(gtx C) D {
											ed := material.Editor(th, &numDiceInput, "1")
											gtx.Constraints.Min.X = gtx.Dp(50)
											gtx.Constraints.Max.X = gtx.Dp(50)
											return text_border.Layout(gtx, ed.Layout)
										}),

										layout.Rigid(layout.Spacer{Width: 20}.Layout),
										layout.Rigid(func(gtx C) D {
											gtx.Constraints.Min.X = gtx.Dp(100)
											gtx.Constraints.Max.X = gtx.Dp(300)
											return material.Body2(th, "Number Faces on Dice").Layout(gtx)
										}),

										layout.Rigid(layout.Spacer{Width: 5}.Layout),
										layout.Rigid(func(gtx C) D {
											ed := material.Editor(th, &typeDiceInput, "20")
											gtx.Constraints.Min.X = gtx.Dp(50)
											gtx.Constraints.Max.X = gtx.Dp(50)
											return text_border.Layout(gtx, ed.Layout)
										}),

										layout.Rigid(layout.Spacer{Width: 20}.Layout),
										layout.Rigid(func(gtx C) D {
											gtx.Constraints.Min.X = gtx.Dp(100)
											gtx.Constraints.Max.X = gtx.Dp(300)
											return material.Body2(th, "+/- to dice roll").Layout(gtx)
										}),

										layout.Rigid(layout.Spacer{Width: 5}.Layout),
										layout.Rigid(func(gtx C) D {
											ed := material.Editor(th, &addRollInput, "0")
											gtx.Constraints.Min.X = gtx.Dp(50)
											gtx.Constraints.Max.X = gtx.Dp(50)
											return text_border.Layout(gtx, ed.Layout)
										}),
									)
								})
							})
						}),

						// Roll Button
						layout.Rigid(func(gtx C) D {
							return roll_margins.Layout(gtx, func(gtx C) D {
								for rollButton.Clicked(gtx) {
									roller.Roller(&numDiceInput, &typeDiceInput, &addRollInput, &Output_editor, &History_editor)
								}
								btn := material.Button(th, &rollButton, "Roll")
								return btn.Layout(gtx)
							})
						}),

						layout.Rigid(layout.Spacer{Height: unit.Dp(25)}.Layout),

						layout.Rigid(func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Dp(200)
							gtx.Constraints.Max.X = gtx.Dp(200)
							return material.Body2(th, "History Output").Layout(gtx)
						}),

						layout.Rigid(func(gtx C) D {
							return output_margins.Layout(gtx, func(gtx C) D {
								return output_border.Layout(gtx, func(gtx C) D {
									gtx.Constraints.Min.Y = gtx.Dp(250)
									gtx.Constraints.Max.Y = gtx.Dp(250)
									return material.Editor(th, &History_editor, "Roll History").Layout(gtx)
								})
							})
						}),
					)
				},
			}

			materialList := material.List(th, &window_list)

			materialList.Layout(gtx, len(listItems), func(gtx C, i int) D {
				return listItems[i](gtx)
			})

			e.Frame(gtx.Ops)
		}
	}
}
