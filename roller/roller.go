package roller

import (
	"math/rand"
	"strconv"

	"gioui.org/widget"
)

func Roller(NumDice, typeDice, addRoll, ed, his *widget.Editor) {
	NumStr := NumDice.Text()
	if NumStr == "" {
		NumStr = "1"
	}
	Num, _ := strconv.Atoi(NumStr)

	DiceStr := typeDice.Text()
	if DiceStr == "" {
		DiceStr = "20"
	}
	Dice, _ := strconv.Atoi(DiceStr)

	BonusStr := addRoll.Text()
	if BonusStr == "" {
		BonusStr = "0"
	}
	Bonus, _ := strconv.Atoi(BonusStr)

	roll_string := "---< Rolled " + NumStr + "d" + DiceStr + " + " + BonusStr + " >---\n"
	ed.SetText(roll_string)
	his.Insert(roll_string)

	var DiceTotalInt int
	ed.SetCaret(ed.Len(), ed.Len())
	for Count := 1; Count <= Num; Count++ {
		CountStr := strconv.Itoa(Count)

		ResultsInt := rand.Intn(Dice) + 1
		ResultsStr := strconv.Itoa(ResultsInt)

		DiceTotalInt = DiceTotalInt + ResultsInt

		ed.SetCaret(ed.Len(), ed.Len())
		insert_string := "Die " + CountStr + ": " + ResultsStr + "\n"
		ed.Insert(insert_string)
		his.Insert(insert_string)
	}

	DiceTotalStr := strconv.Itoa(DiceTotalInt)
	RollTotalStr := strconv.Itoa(DiceTotalInt + Bonus)

	totlal_string := "---< Total of the Dice " + DiceTotalStr + " + Bonus/Penalty " + BonusStr + " = Grand Total " + RollTotalStr + " >---\n"
	ed.Insert(totlal_string)
	his.Insert(totlal_string)

	ed.SetCaret(ed.Len(), ed.Len())
	his.SetCaret(his.Len(), his.Len())
}
