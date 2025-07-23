package roller

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"

	"gioui.org/widget"
)

type RollArgs struct {
	NumDiceEd      *widget.Editor
	TypeDiceEd     *widget.Editor
	DiffEd         *widget.Editor
	TargetNumberEd *widget.Editor
	BonusEd        *widget.Editor
	OutputEd       *widget.Editor
	HistoryEd      *widget.Editor
	DoubleS        [4]int
	RollTillGoneS  [4]int
}

func rollArray(t int, n int) ([]int, error) {
	if t <= 2 {
		return nil, errors.New("rollArray t argument must be >= 2")
	}

	if n <= 0 {
		return nil, errors.New("rollArray n argument must be >= 1")
	}

	var a []int
	for c := 1; c <= n; c++ {
		r := rand.Intn(t) + 1
		a = append(a, r)
	}

	return a, nil
}

// func Roller(numDiceEd, typeDiceEd, bonusEd, outputEd, historyEd *widget.Editor)
func Roller(args *RollArgs) {
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

	roll_string := "---< Rolled " + numStr + "d" + diceStr + " + " + bonusStr + " >---\n"
	args.OutputEd.SetText(roll_string)
	args.HistoryEd.Insert(roll_string)

	var diceTotalInt int
	//outputEd.SetCaret(outputEd.Len(), outputEd.Len())
	for count := 1; count <= num; count++ {
		countStr := strconv.Itoa(count)

		resultInt := rand.Intn(dice) + 1
		resultStr := strconv.Itoa(resultInt)

		diceTotalInt = diceTotalInt + resultInt

		args.OutputEd.SetCaret(args.OutputEd.Len(), args.OutputEd.Len())
		insert_string := "Die " + countStr + ": " + resultStr + "\n"
		args.OutputEd.Insert(insert_string)
		args.HistoryEd.Insert(insert_string)
	}

	diceTotalStr := strconv.Itoa(diceTotalInt)
	rollTotalStr := strconv.Itoa(diceTotalInt + bonus)

	total_string := "---< Total of the Dice " + diceTotalStr + " + Bonus/Penalty " + bonusStr + " = Grand Total " + rollTotalStr + " >---\n"
	args.OutputEd.Insert(total_string)
	args.HistoryEd.Insert(total_string)

	args.OutputEd.SetCaret(args.OutputEd.Len(), args.OutputEd.Len())
	args.HistoryEd.SetCaret(args.HistoryEd.Len(), args.HistoryEd.Len())
}

func StorytellerSystem(args *RollArgs) {
	diceStr := args.TypeDiceEd.Text()
	if diceStr == "" {
		diceStr = "10"
	}
	diceInt, _ := strconv.Atoi(diceStr)

	numStr := args.NumDiceEd.Text()
	num, _ := strconv.Atoi(numStr)
	if num < 1 {
		args.OutputEd.SetText("---< Please enter a non-zero number of dice? >---")
		return
	}

	diffStr := args.DiffEd.Text()
	if diffStr == "" {
		diffStr = "3"
	}
	diff, _ := strconv.Atoi(diffStr)

	targetNumStr := args.TargetNumberEd.Text()
	if targetNumStr == "" {
		targetNumStr = "7"
	}
	targetNum, _ := strconv.Atoi(targetNumStr)

	var rerollStr string
	var sliceStr []string
	for _, n := range args.RollTillGoneS {
		if n != 0 {
			sliceStr = append(sliceStr, strconv.Itoa(n))
		}
	}
	if len(sliceStr) == 0 {
		rerollStr = "None,"
	} else {
		rerollStr = strings.Join(sliceStr, ",")
	}

	var doubleStr string
	sliceStr = sliceStr[:0]
	for _, n := range args.DoubleS {
		if n != 0 {
			sliceStr = append(sliceStr, strconv.Itoa(n))
		}
	}
	if len(sliceStr) == 0 {
		doubleStr = "None,"
	} else {
		doubleStr = strings.Join(sliceStr, ",")
	}

	roll_string := "---< Rolled " + numStr + "d" + diceStr + " at Diffculty " + targetNumStr + " >---\n"
	details_string := "---< Details: Target Number " + targetNumStr + ", Reroll " + rerollStr + " and Doubles " + doubleStr + " >---\n"
	args.OutputEd.SetText(roll_string)
	args.HistoryEd.Insert(roll_string)
	args.HistoryEd.Insert(details_string)
	args.OutputEd.SetCaret(args.OutputEd.Len(), args.OutputEd.Len())

	var suxTotalInt int
	rollExalted(diceInt, num, targetNum, &suxTotalInt, args)

	suxTotalStr := strconv.Itoa(suxTotalInt)

	var total_string string
	// think about switch case statment instead of if statements
	if suxTotalInt-diff >= 1 {
		thresholdStr := strconv.Itoa(suxTotalInt - diff)
		total_string = "---< " + suxTotalStr + " Success with " + thresholdStr + " Threshold Successes >---\n"
	}
	if suxTotalInt-diff == 0 {
		total_string = "---< " + suxTotalStr + " Success >---\n"
	}
	if suxTotalInt-diff < 0 {
		failStr := strconv.Itoa((suxTotalInt - diff) * -1)
		total_string = "---< " + suxTotalStr + " Success, Failed by " + failStr + " Successes >---\n"
	}

	args.OutputEd.Insert(total_string)
	args.HistoryEd.Insert(total_string)

	args.OutputEd.SetCaret(args.OutputEd.Len(), args.OutputEd.Len())
	args.HistoryEd.SetCaret(args.HistoryEd.Len(), args.HistoryEd.Len())
}

func rollExalted(diceInt, num, targetNum int, suxTotalInt *int, args *RollArgs) {
	results, err := rollArray(diceInt, num)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var sliceStr []string
	for _, n := range results {
		sliceStr = append(sliceStr, strconv.Itoa(n))
	}
	resultsStr := strings.Join(sliceStr, ",") + "\n"
	args.OutputEd.Insert(resultsStr)
	args.HistoryEd.Insert(resultsStr)

	var rerollDice []int
	for _, result := range results {
		if slices.Contains(args.RollTillGoneS[:], result) {
			rerollDice = append(rerollDice, result)
		}

		if slices.Contains(args.DoubleS[:], result) {
			(*suxTotalInt)++
		}

		if result >= targetNum {
			(*suxTotalInt)++
		}
	}

	if len(rerollDice) != 0 {
		var sliceStr []string
		for _, n := range rerollDice {
			if n != 0 {
				sliceStr = append(sliceStr, strconv.Itoa(n))
			}
		}
		rerollDiceStr := "Rerolling " + strings.Join(sliceStr, ",") + "\n"
		args.OutputEd.Insert(rerollDiceStr)
		args.HistoryEd.Insert(rerollDiceStr)

		rerollInt := len(rerollDice)
		rollExalted(diceInt, rerollInt, targetNum, suxTotalInt, args)
	}
}
