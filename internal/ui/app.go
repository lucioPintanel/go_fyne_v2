package ui

import (
	"crud/internal/models"
	"crud/internal/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

/* func SetupUI(w fyne.Window) {
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
} */

type ProductUI struct {
	ProductListBinding binding.StringList
}

func NewProductUI() *ProductUI {
	ui := &ProductUI{
		ProductListBinding: binding.NewStringList(),
	}
	ui.loadProductsFromDB() // Carrega os produtos do banco de dados ao iniciar
	return ui
}

func (ui *ProductUI) SetupUI(w fyne.Window) {
	productList := widget.NewListWithData(
		ui.ProductListBinding,
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(item binding.DataItem, o fyne.CanvasObject) {
			str, _ := item.(binding.String).Get()
			o.(*widget.Label).SetText(str)
		},
	)
	productList.OnSelected = func(id widget.ListItemID) {
		ui.showEditForm(id)
		productList.Unselect(id)
	}
	btnAdd := widget.NewButton("Adicionar", func() { ui.showForm() })
	content := container.NewBorder(nil, container.NewHBox(btnAdd), nil, nil, container.NewVScroll(productList))
	w.SetContent(content)
}

func (ui *ProductUI) loadProductsFromDB() {
	products_serv, _ := services.GetAllProducts()
	var products []string
	for _, p := range products_serv {
		products = append(products, p.Description)
	}
	ui.ProductListBinding.Set(products)
}

func (ui *ProductUI) showForm() {
	form := widget.NewForm(
		widget.NewFormItem("Description", widget.NewEntry()),
		widget.NewFormItem("Product Type", widget.NewEntry()),
	)
	formWindow := fyne.CurrentApp().NewWindow("Add Product")
	saveButton := widget.NewButton("Save", func() {
		description := form.Items[0].Widget.(*widget.Entry).Text
		productType := form.Items[1].Widget.(*widget.Entry).Text
		product := models.Product{
			Description: description,
			ProductType: productType,
		}
		services.CreateProduct(product)
		ui.ProductListBinding.Append(description)
		formWindow.Close()
		ui.loadProductsFromDB()
	})
	formWindow.SetContent(container.NewVBox(form, saveButton))
	formWindow.Resize(fyne.NewSize(300, 150))
	formWindow.Show()
}

func (ui *ProductUI) showEditForm(id widget.ListItemID) {
	item, _ := ui.ProductListBinding.GetItem(id)
	str, _ := item.(binding.String).Get()
	// Retrieve the product from your data source based on 'str' or 'id'
	product := getProductFromDB(str)
	if product.Description == "" {
		return
	}

	descriptionEntry := widget.NewEntry()
	descriptionEntry.SetText(product.Description)
	productTypeEntry := widget.NewEntry()
	productTypeEntry.SetText(product.ProductType)

	form := widget.NewForm(
		widget.NewFormItem("Description", descriptionEntry),
		widget.NewFormItem("Product Type", productTypeEntry),
	)
	formWindow := fyne.CurrentApp().NewWindow("Edit Product")
	saveButton := widget.NewButton("Save", func() {
		product.Description = descriptionEntry.Text
		product.ProductType = productTypeEntry.Text
		services.UpdateProduct(product)
		ui.ProductListBinding.SetValue(id, product.Description)
		formWindow.Close()
		ui.loadProductsFromDB()
	})
	formWindow.SetContent(container.NewVBox(form, saveButton))
	formWindow.Resize(fyne.NewSize(300, 150))
	formWindow.Show()
}

func getProductFromDB(description string) models.Product {
	// Implement your logic to retrieve the product from the database
	// For example:
	product, err := services.GetProductByDescription(description)
	if err != nil {
		return models.Product{}
	}
	return product
}
