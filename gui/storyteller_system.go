package gui

import (
	"go-roll/roller"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var checkDouble10s = widget.Bool{
	Value: true,
}

var (
	checkDouble9s widget.Bool
	checkDouble8s widget.Bool
	checkDouble7s widget.Bool
)

var (
	reroll10s widget.Bool
	reroll6s  widget.Bool
	reroll5s  widget.Bool
	reroll1s  widget.Bool
)

var (
	doubleS [4]int
	rerollS [4]int
)

var storytellerList = []layout.Widget{
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
						// TODO fix hint
						return material.Editor(th, &Output_editor, "Enter a number of dice,\nNumber of sides of dice,\nand +/- to the dice roll.").Layout(gtx)
					})
				})
			}),

			layout.Rigid(func(gtx C) D {
				return layout.Center.Layout(gtx, func(gtx C) D {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Dp(50)
							gtx.Constraints.Max.X = gtx.Dp(300)
							return material.Body2(th, "Target Number").Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Width: 10}.Layout),
						layout.Rigid(func(gtx C) D {
							ed := material.Editor(th, &TargetNumberInput, "7")
							gtx.Constraints.Min.X = gtx.Dp(50)
							gtx.Constraints.Max.X = gtx.Dp(50)
							return text_border.Layout(gtx, ed.Layout)
						}),
					)
				})
			}),

			layout.Rigid(layout.Spacer{Height: 25}.Layout),

			layout.Rigid(func(gtx C) D {
				return layout.Center.Layout(gtx, func(gtx C) D {
					// return text_margins.Layout(gtx, func(gtx C) D {
					// Container for the Columns
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(gtx,

						// Column 1
						layout.Rigid(func(gtx C) D {
							return layout.Center.Layout(gtx, func(gtx C) D {
								return layout.Flex{
									Axis: layout.Vertical,
								}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.CheckBox(th, &reroll10s, "Reroll 10s").Layout(gtx)
									}),

									layout.Rigid(layout.Spacer{Height: unit.Dp(5)}.Layout),

									layout.Rigid(func(gtx C) D {
										return material.CheckBox(th, &reroll6s, "Reroll 6s").Layout(gtx)
									}),

									layout.Rigid(layout.Spacer{Height: unit.Dp(5)}.Layout),

									layout.Rigid(func(gtx C) D {
										return material.CheckBox(th, &reroll5s, "Reroll 5s").Layout(gtx)
									}),

									layout.Rigid(layout.Spacer{Height: unit.Dp(5)}.Layout),

									layout.Rigid(func(gtx C) D {
										return material.CheckBox(th, &reroll1s, "Reroll 1s").Layout(gtx)
									}),
								)
							})
						}),

						layout.Rigid(layout.Spacer{Width: unit.Dp(100)}.Layout),

						// Column 2
						layout.Rigid(func(gtx C) D {
							return layout.Center.Layout(gtx, func(gtx C) D {
								return layout.Flex{
									Axis: layout.Vertical,
								}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.CheckBox(th, &checkDouble10s, "Double 10s").Layout(gtx)
									}),

									layout.Rigid(layout.Spacer{Height: unit.Dp(5)}.Layout),

									layout.Rigid(func(gtx C) D {
										return material.CheckBox(th, &checkDouble9s, "Double 9s").Layout(gtx)
									}),

									layout.Rigid(layout.Spacer{Height: unit.Dp(5)}.Layout),

									layout.Rigid(func(gtx C) D {
										return material.CheckBox(th, &checkDouble8s, "Double 8s").Layout(gtx)
									}),

									layout.Rigid(layout.Spacer{Height: unit.Dp(5)}.Layout),

									layout.Rigid(func(gtx C) D {
										return material.CheckBox(th, &checkDouble7s, "Double 7s").Layout(gtx)
									}),
								)
							})
						}),
					)
					// })
				})
			}),

			layout.Rigid(layout.Spacer{Height: 25}.Layout),

			// Rigid Box of the Horizontal Text Inputs
			layout.Rigid(func(gtx C) D {
				return layout.Center.Layout(gtx, func(gtx C) D {
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
							ed := material.Editor(th, &numDiceInput, "0")
							gtx.Constraints.Min.X = gtx.Dp(50)
							gtx.Constraints.Max.X = gtx.Dp(50)
							return text_border.Layout(gtx, ed.Layout)
						}),

						layout.Rigid(layout.Spacer{Width: 20}.Layout),

						layout.Rigid(func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Dp(50)
							gtx.Constraints.Max.X = gtx.Dp(300)
							return material.Body2(th, "Diffculty").Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Width: 5}.Layout),
						layout.Rigid(func(gtx C) D {
							ed := material.Editor(th, &diffInput, "3")
							gtx.Constraints.Min.X = gtx.Dp(50)
							gtx.Constraints.Max.X = gtx.Dp(50)
							return text_border.Layout(gtx, ed.Layout)
						}),
					)
				})
			}),

			layout.Rigid(layout.Spacer{Height: 50}.Layout),

			// Roll Button
			layout.Rigid(func(gtx C) D {
				return roll_margins.Layout(gtx, func(gtx C) D {
					for rollButton.Clicked(gtx) {
						checkBoxes()
						args := roller.RollArgs{
							NumDiceEd:      &numDiceInput,
							TypeDiceEd:     &typeDiceInput,
							DiffEd:         &diffInput,
							TargetNumberEd: &TargetNumberInput,
							OutputEd:       &Output_editor,
							HistoryEd:      &History_editor,
							DoubleS:        doubleS,
							RollTillGoneS:  rerollS,
						}
						roller.StorytellerSystem(&args)
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

func checkBoxes() {
	// checking rerolls
	if reroll10s.Value {
		rerollS[0] = 10
	} else {
		rerollS[0] = 0
	}
	if reroll6s.Value {
		rerollS[1] = 6
	} else {
		rerollS[1] = 0
	}
	if reroll5s.Value {
		rerollS[2] = 5
	} else {
		rerollS[2] = 0
	}
	if reroll1s.Value {
		rerollS[3] = 1
	} else {
		rerollS[3] = 0
	}

	// checking doubles
	if checkDouble10s.Value {
		doubleS[0] = 10
	} else {
		doubleS[0] = 0
	}
	if checkDouble9s.Value {
		doubleS[1] = 9
	} else {
		doubleS[1] = 0
	}
	if checkDouble8s.Value {
		doubleS[2] = 8
	} else {
		doubleS[2] = 0
	}
	if checkDouble7s.Value {
		doubleS[3] = 7
	} else {
		doubleS[3] = 0
	}
}
