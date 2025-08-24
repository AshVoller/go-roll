package roller

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO Handle errors from Atoi lines
// func BasicRoller(numDiceEd, typeDiceEd, bonusEd, outputEd, historyEd *widget.Editor)
func BasicRoller(args *RollArgs) {
	numStr := args.NumDiceEd.Text()
	if numStr == "" {
		numStr = "1"
	}
	num, _ := strconv.Atoi(numStr)

	diceStr := args.TypeDiceEd.Text()
	if diceStr == "" {
		diceStr = "20"
	}
	dice, _ := strconv.Atoi(diceStr)

	addDieStr := args.AddDieEd.Text()
	if addDieStr == "" {
		addDieStr = "0"
	}
	bonusDie, _ := strconv.Atoi(addDieStr)

	addTotalStr := args.AddTotalEd.Text()
	if addTotalStr == "" {
		addTotalStr = "0"
	}
	bonusTotal, _ := strconv.Atoi(addTotalStr)

	roll_string := fmt.Sprintf("---< Rolled %sd%s + %s >---\n", numStr, diceStr, addTotalStr)
	args.OutputEd.SetText(roll_string)
	args.HistoryEd.Insert(roll_string)

	results, err := rollArray(dice, num)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var diceTotalInt int
	var sliceStr []string
	for _, n := range results {
		nDie := n + bonusDie
		diceTotalInt = diceTotalInt + nDie
		var str string
		if addDieStr == "0" || addDieStr == "" {
			str = fmt.Sprintf("%d", n)
		} else {
			str = fmt.Sprintf("%d", nDie)
		}
		sliceStr = append(sliceStr, str)
	}

	args.OutputEd.SetCaret(args.OutputEd.Len(), args.OutputEd.Len())
	resultsStr := strings.Join(sliceStr, ",") + "\n"
	args.OutputEd.Insert(resultsStr)
	args.HistoryEd.Insert(resultsStr)

	diceTotalStr := strconv.Itoa(diceTotalInt)
	rollTotalStr := strconv.Itoa(diceTotalInt + bonusTotal)

	total_string := fmt.Sprintf("---< Total of the Dice %s + Bonus/Penalty %s = Grand Total %s >---\n", diceTotalStr, addTotalStr, rollTotalStr)
	args.OutputEd.Insert(total_string)
	args.HistoryEd.Insert(total_string)

	args.OutputEd.SetCaret(args.OutputEd.Len(), args.OutputEd.Len())
	args.HistoryEd.SetCaret(args.HistoryEd.Len(), args.HistoryEd.Len())
}
