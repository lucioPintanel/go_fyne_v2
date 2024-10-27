package services

import (
	"crud/internal/database"
	"crud/internal/models"
	"crud/internal/repository"
	"database/sql"
)

func GetAllProducts() ([]models.Product, error) {
	return repository.GetAllProducts(database.DB)
}

func GetProductByDescription(description string) (models.Product, error) {
	return repository.GetProductByDescription(database.DB, description)
}

func CreateProduct(product models.Product) error {
	return repository.CreateProduct(database.DB, product)
}

func UpdateProduct(product models.Product) error {
	return repository.UpdateProduct(database.DB, product)
}

func DeleteProduct(id int) (sql.Result, error) {
	return repository.DeleteProduct(database.DB, id)
}
