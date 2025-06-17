package main

import (
	//    "fmt"
	//    "math/rand"
	//    "strconv"
	//    "image"
	//	"image/color"
	"log"
	"os"

	"gioui.org/app"
	//	"gioui.org/op"

	//    "gioui.org/op/clip"
	//	"gioui.org/text"
	"gioui.org/unit"

	//    "gioui.org/op/paint"
	//	"gioui.org/layout"
	//	"gioui.org/widget"
	//	"gioui.org/widget/material"
	// "gioui.org/io/system"

	"diceroller/gui"
)

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Diceroller"))
		w.Option(app.Size(unit.Dp(1000), unit.Dp(1000)))

		if err := gui.Gui(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}
