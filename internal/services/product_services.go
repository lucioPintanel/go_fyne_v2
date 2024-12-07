package services

import (
	"crud/internal/database"
	"crud/internal/models"
	"crud/internal/repository"
	"database/sql"
)

var Repo repository.Repository = repository.Repo

func GetAllProducts() ([]models.Product, error) {
	return Repo.GetAllProducts(database.DB)
}

func GetProductByDescription(description string) (models.Product, error) {
	return Repo.GetProductByDescription(database.DB, description)
}

func CreateProduct(product models.Product) error {
	return Repo.CreateProduct(database.DB, product)
}

func UpdateProduct(product models.Product) error {
	return Repo.UpdateProduct(database.DB, product)
}

func DeleteProduct(id int) (sql.Result, error) {
	return Repo.DeleteProduct(database.DB, id)
}
