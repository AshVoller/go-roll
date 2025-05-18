package main

import (
	//    "fmt"
	//    "math/rand"
	//    "strconv"
	//    "image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"

	//    "gioui.org/op/clip"
	"gioui.org/text"
	"gioui.org/unit"

	//    "gioui.org/op/paint"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	// "gioui.org/io/system"
)

type C = layout.Context
type D = layout.Dimensions

var text_margins = layout.Inset{
	Top:    unit.Dp(40),
	Bottom: unit.Dp(40),
	Right:  unit.Dp(170),
	Left:   unit.Dp(170),
}

var text_border = widget.Border{
	Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
	CornerRadius: unit.Dp(3),
	Width:        unit.Dp(2),
}

var output_margins = layout.Inset{
	Top:    unit.Dp(40),
	Bottom: unit.Dp(40),
	Right:  unit.Dp(170),
	Left:   unit.Dp(170),
}

var output_border = widget.Border{
	Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
	CornerRadius: unit.Dp(3),
	Width:        unit.Dp(2),
}

var output_editor = widget.Editor{
	ReadOnly: true,
}

var history_editor = widget.Editor{
	ReadOnly: true,
}

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Diceroller"))
		w.Option(app.Size(unit.Dp(1000), unit.Dp(1000)))

		if err := gui(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}

func gui(w *app.Window) error {
	// Roll Button
	var rollButton widget.Clickable

	// Button Size
	// inputTextSize := clip.Rect(image.Rect(0, 0, 100, 100)).Push(ops)

	// Maybe Turn into a struct ?
	// Number of Dice
	var numDiceInput widget.Editor
	numDiceInput.SingleLine = true
	numDiceInput.Alignment = text.Middle

	// Type of Dice
	var typeDiceInput widget.Editor
	typeDiceInput.SingleLine = true
	typeDiceInput.Alignment = text.Middle

	// +/- to Dice Roll
	var addRollInput widget.Editor
	addRollInput.SingleLine = true
	addRollInput.Alignment = text.Middle

	output_editor.SetText("Output\nOutput\nOutput\nOutput\nOutput\nOutput\nOutput\nOutput\nOutput\nOutput\nOutput\nOutput")

	history_editor.SetText("History\nHistory\nHistory\nHistory\nHistory\nHistory\nHistory\nHistory\nHistory\nHistory\nHistory")

	var ops op.Ops

	th := material.NewTheme()

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Flex Box of the whole page
			layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Start,
				//Spacing:   layout.SpaceStart, // come back to later
			}.Layout(gtx,
				layout.Rigid(layout.Spacer{Height: unit.Dp(25)}.Layout),

				layout.Rigid(func(gtx C) D {
					return output_margins.Layout(gtx, func(gtx C) D {
						return output_border.Layout(gtx, func(gtx C) D {
							return material.Editor(th, &output_editor, "").Layout(gtx)
						})
					})
				}),

				// Rigid Box of the Horizontal Text Inputs
				layout.Rigid(func(gtx C) D {
					return text_margins.Layout(gtx, func(gtx C) D {
						return layout.Flex{
							Axis: layout.Horizontal,
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
							layout.Rigid(layout.Spacer{Width: 20}.Layout),
						)
					})
				}),

				// Roll Button
				layout.Rigid(func(gtx C) D {
					margins := layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(55),
						Left:   unit.Dp(55),
					}

					return margins.Layout(gtx, func(gtx C) D {
						for rollButton.Clicked(gtx) {
							rolled()
						}
						btn := material.Button(th, &rollButton, "Roll")
						return btn.Layout(gtx)
					})
				}),

				layout.Rigid(layout.Spacer{Height: unit.Dp(25)}.Layout),

				layout.Rigid(func(gtx C) D {
					return output_margins.Layout(gtx, func(gtx C) D {
						return output_border.Layout(gtx, func(gtx C) D {
							return material.Editor(th, &history_editor, "").Layout(gtx)
						})
					})
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}

func AppendEditor(ed widget.Editor, S string) {
	current_text := ed.Text()
	new_text := current_text + S
	ed.SetText(new_text)
}

func rolled() {
	output_editor.SetText("")
	var output_text string = "button clicked"
	AppendEditor(output_editor, output_text)
}
