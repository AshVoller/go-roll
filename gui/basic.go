package gui

import (
	"go-roll/roller"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

var basicList = []layout.Widget{
	func(gtx C) D {
		// Flex Box of the whole page
		return layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
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
						args := roller.RollArgs{
							NumDiceEd:  &numDiceInput,
							TypeDiceEd: &typeDiceInput,
							BonusEd:    &addRollInput,
							OutputEd:   &Output_editor,
							HistoryEd:  &History_editor,
						}
						roller.Roller(&args)
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
