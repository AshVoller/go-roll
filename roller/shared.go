package roller

import (
	"errors"
	"math/rand"

	"gioui.org/widget"
)

type RollArgs struct {
	NumDiceEd      *widget.Editor
	TypeDiceEd     *widget.Editor
	DiffEd         *widget.Editor
	TargetNumberEd *widget.Editor
	AddTotalEd     *widget.Editor
	AddDieEd       *widget.Editor
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
