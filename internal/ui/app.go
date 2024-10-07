package ui

import (
	"exemplo.com/crud/internal/models"
	"exemplo.com/crud/internal/services"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SetupUI(w fyne.Window) {
	productList := widget.NewList(
		func() int {
			products, _ := services.GetAllProducts()
			return len(products)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			products, _ := services.GetAllProducts()
			o.(*widget.Label).SetText(products[i].Description)
		},
	)

	form := widget.NewForm(
		widget.NewFormItem("Description", widget.NewEntry()),
		widget.NewFormItem("Product Type", widget.NewEntry()),
	)

	saveButton := widget.NewButton("Save", func() {
		description := form.Items[0].Widget.(*widget.Entry).Text
		productType := form.Items[1].Widget.(*widget.Entry).Text
		product := models.Product{Description: description, ProductType: productType}
		services.CreateProduct(product)
		productList.Refresh()
	})

	content := container.NewVBox(productList, form, saveButton)
	w.SetContent(content)
}
