package main

import (
	"crud/internal/database"
	"crud/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	w := a.NewWindow("Product App")

	database.InitDB()

	lui := ui.NewProductUI()

	lui.SetupUI(w)

	w.Resize(fyne.NewSize(400, 400))

	w.ShowAndRun()
}
