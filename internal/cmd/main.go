package main

import (
	"exemplo.com/crud/internal/database"
	"exemplo.com/crud/internal/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Product App")

	database.InitDB()

	ui.SetupUI(w)

	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}
