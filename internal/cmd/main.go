package main

import (
	"exemplo.com/crud/internal/ui"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Product App")

	ui.SetupUI(w)

	w.ShowAndRun()
}
