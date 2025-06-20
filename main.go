package main

import (
	"log"
	"os"

	"gioui.org/app"

	"go-roll/gui"
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
