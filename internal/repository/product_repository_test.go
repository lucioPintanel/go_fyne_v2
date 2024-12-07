package repository

import (
	"testing"
	"time"

	"crud/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "description", "product_type", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "Product 1", "Type 1", time.Now(), time.Now(), nil).
		AddRow(2, "Product 2", "Type 2", time.Now(), time.Now(), nil)

	mock.ExpectQuery("SELECT id, description, product_type, created_at, updated_at, deleted_at FROM products WHERE deleted_at IS NULL").
		WillReturnRows(rows)

	repo := repositoryImpl{}
	products, err := repo.GetAllProducts(db)
	assert.NoError(t, err)
	assert.Len(t, products, 2)
}

func TestGetProductByDescription(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "description", "product_type", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "Product 1", "Type 1", time.Now(), time.Now(), nil)

	mock.ExpectQuery(`SELECT id, description, product_type, created_at, updated_at, deleted_at FROM products WHERE description = \? AND deleted_at IS NULL`).
		WithArgs("Product 1").
		WillReturnRows(row)

	repo := repositoryImpl{}
	product, err := repo.GetProductByDescription(db, "Product 1")
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", product.Description)
	assert.Equal(t, "Type 1", product.ProductType)
}

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product := models.Product{
		Description: "New Product",
		ProductType: "New Type",
	}

	mock.ExpectExec("INSERT INTO products").
		WithArgs(product.Description, product.ProductType).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repositoryImpl{}
	err = repo.CreateProduct(db, product)
	assert.NoError(t, err)
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product := models.Product{
		ID:          1,
		Description: "Updated Product",
		ProductType: "Updated Type",
	}

	mock.ExpectExec(`UPDATE products SET description = \?, product_type = \?, updated_at = \? WHERE id = \?`).
		WithArgs(product.Description, product.ProductType, sqlmock.AnyArg(), product.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repositoryImpl{}
	err = repo.UpdateProduct(db, product)
	assert.NoError(t, err)
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	id := 1

	mock.ExpectExec(`UPDATE products SET deleted_at = \? WHERE id = \?`).
		WithArgs(sqlmock.AnyArg(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repositoryImpl{}
	result, err := repo.DeleteProduct(db, id)
	assert.NoError(t, err)

	rowsAffected, err := result.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowsAffected)
}
