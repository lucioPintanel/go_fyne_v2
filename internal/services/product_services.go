package services

import (
	"exemplo.com/crud/internal/database"
	"exemplo.com/crud/internal/models"
	"exemplo.com/crud/internal/repository"
)

func GetAllProducts() ([]models.Product, error) {
	return repository.GetAllProducts(database.DB)
}

func CreateProduct(product models.Product) error {
	return repository.CreateProduct(database.DB, product)
}

func UpdateProduct(product models.Product) error {
	return repository.UpdateProduct(database.DB, product)
}

func DeleteProduct(id int) error {
	return repository.DeleteProduct(database.DB, id)
}
