package roller

import (
	//    "fmt"
	//    "math/rand"
	//    "strconv"
	//    "image"
	//	"image/color"
	//	"log"
	//	"os"

	//	"gioui.org/app"
	//	"gioui.org/op"

	//    "gioui.org/op/clip"
	//	"gioui.org/text"
	//	"gioui.org/unit"

	//    "gioui.org/op/paint"
	//	"gioui.org/layout"

	"math/rand"
	"strconv"

	"gioui.org/widget"
	//	"gioui.org/widget/material"
	// "gioui.org/io/system"
)

func Roller(NumDice, typeDice, addRoll, ed *widget.Editor) {
	NumStr := NumDice.Text()
	Num, _ := strconv.Atoi(NumStr)

	DiceStr := typeDice.Text()
	Dice, _ := strconv.Atoi(DiceStr)

	BonusStr := addRoll.Text()
	Bonus, _ := strconv.Atoi(BonusStr)

	ed.SetCaret(ed.Len(), ed.Len())
	roll_string := "---< Rolled " + NumStr + "d" + DiceStr + " + " + BonusStr + " >---\n"
	ed.Insert(roll_string)

	var DiceTotalInt int

	for Count := 1; Count <= Num; Count++ {
		CountStr := strconv.Itoa(Count)

		ResultsInt := rand.Intn(Dice) + 1
		ResultsStr := strconv.Itoa(ResultsInt)

		DiceTotalInt = DiceTotalInt + ResultsInt

		ed.SetCaret(ed.Len(), ed.Len())
		insert_string := "Die " + CountStr + ": " + ResultsStr + "\n"
		ed.Insert(insert_string)
		ed.SetCaret(ed.Len(), ed.Len())
		// fmt.Printf("Die %v: %v\n", Count, ResultsStr)
	}

	DiceTotalStr := strconv.Itoa(DiceTotalInt)
	RollTotalStr := strconv.Itoa(DiceTotalInt + Bonus)

	ed.SetCaret(ed.Len(), ed.Len())
	totlal_string := "---< Dice " + DiceTotalStr + " + Bonus/Penalty" + BonusStr + " = Total " + RollTotalStr + " >---\n"
	ed.Insert(totlal_string)
}
