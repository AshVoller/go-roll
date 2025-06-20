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
		w.Option(app.Size(1000, 1000))
		w.Option(app.MinSize(900, 500))

		if err := gui.Gui(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}
