package ui

import (
	"crud/internal/models"
	"crud/internal/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

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
	delButton := widget.NewButton("Delete", func() {
		services.DeleteProduct(product.ID)
		err := removeItemFromStringList(ui.ProductListBinding, id)
		if err != nil {
			return
		}
		//ui.ProductListBinding.Set()
		formWindow.Close()
		ui.loadProductsFromDB()
	})
	formWindow.SetContent(container.NewVBox(form, container.NewVBox(saveButton, delButton)))
	formWindow.Resize(fyne.NewSize(300, 150))
	formWindow.Show()
}

func getProductFromDB(description string) models.Product {
	product, err := services.GetProductByDescription(description)
	if err != nil {
		return models.Product{}
	}
	return product
}

func removeItemFromStringList(list binding.StringList, index int) error {
	items, err := list.Get()
	if err != nil {
		return err
	}

	// Remover o item do slice
	items = append(items[:index], items[index+1:]...)

	// Atualizar o binding com o novo slice
	return list.Set(items)
}
