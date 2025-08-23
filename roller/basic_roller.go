package roller

import (
	"fmt"
	"strconv"
	"strings"
)

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

	bonusStr := args.BonusEd.Text()
	if bonusStr == "" {
		bonusStr = "0"
	}
	bonus, _ := strconv.Atoi(bonusStr)

	roll_string := fmt.Sprintf("---< Rolled %sd%s + %s >---\n", numStr, diceStr, bonusStr)
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
		diceTotalInt = diceTotalInt + n
		sliceStr = append(sliceStr, strconv.Itoa(n))
	}

	args.OutputEd.SetCaret(args.OutputEd.Len(), args.OutputEd.Len())
	resultsStr := strings.Join(sliceStr, ",") + "\n"
	args.OutputEd.Insert(resultsStr)
	args.HistoryEd.Insert(resultsStr)

	diceTotalStr := strconv.Itoa(diceTotalInt)
	rollTotalStr := strconv.Itoa(diceTotalInt + bonus)

	total_string := fmt.Sprintf("---< Total of the Dice %s + Bonus/Penalty %s = Grand Total %s >---\n", diceTotalStr, bonusStr, rollTotalStr)
	args.OutputEd.Insert(total_string)
	args.HistoryEd.Insert(total_string)

	args.OutputEd.SetCaret(args.OutputEd.Len(), args.OutputEd.Len())
	args.HistoryEd.SetCaret(args.HistoryEd.Len(), args.HistoryEd.Len())
}
