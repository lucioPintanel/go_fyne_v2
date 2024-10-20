package repository

import (
	"database/sql"
	"time"

	"crud/internal/models"
)

func GetAllProducts(db *sql.DB) ([]models.Product, error) {
	rows, err := db.Query("SELECT id, description, product_type, created_at, updated_at, deleted_at FROM products WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var deletedAt sql.NullTime
		err := rows.Scan(&product.ID, &product.Description, &product.ProductType, &product.CreatedAt, &product.UpdatedAt, &deletedAt)
		if err != nil {
			return nil, err
		}
		if deletedAt.Valid {
			product.DeletedAt = &deletedAt.Time
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProductByDescription(db *sql.DB, description string) (models.Product, error) {
	var product models.Product
	var deletedAt sql.NullTime

	query := "SELECT id, description, product_type, created_at, updated_at, deleted_at FROM products WHERE description = ? AND deleted_at IS NULL"
	err := db.QueryRow(query, description).Scan(&product.ID, &product.Description, &product.ProductType, &product.CreatedAt, &product.UpdatedAt, &deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, nil // Retorna produto vazio se não for encontrado
		}
		return product, err
	}

	if deletedAt.Valid {
		product.DeletedAt = &deletedAt.Time
	}
	return product, nil
}

func CreateProduct(db *sql.DB, product models.Product) error {
	_, err := db.Exec("INSERT INTO products (description, product_type) VALUES (?, ?)", product.Description, product.ProductType)
	return err
}

func UpdateProduct(db *sql.DB, product models.Product) error {
	_, err := db.Exec("UPDATE products SET description = ?, product_type = ?, updated_at = ? WHERE id = ?", product.Description, product.ProductType, time.Now(), product.ID)
	return err
}

func DeleteProduct(db *sql.DB, id int) error {
	_, err := db.Exec("UPDATE products SET deleted_at = ? WHERE id = ?", time.Now(), id)
	return err
}
