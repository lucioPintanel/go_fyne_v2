package services

import (
	"crud/internal/models"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository é o mock do repositório
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAllProducts(db *sql.DB) ([]models.Product, error) {
	args := m.Called(db)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockRepository) GetProductByDescription(db *sql.DB, description string) (models.Product, error) {
	args := m.Called(db, description)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockRepository) CreateProduct(db *sql.DB, product models.Product) error {
	args := m.Called(db, product)
	return args.Error(0)
}

func (m *MockRepository) UpdateProduct(db *sql.DB, product models.Product) error {
	args := m.Called(db, product)
	return args.Error(0)
}

func (m *MockRepository) DeleteProduct(db *sql.DB, id int) (sql.Result, error) {
	args := m.Called(db, id)
	return args.Get(0).(sql.Result), args.Error(1)
}

var (
	db       *sql.DB
	mockRepo *MockRepository
)

func setup() {
	mockRepo = new(MockRepository)
	Repo = mockRepo
}

func TestGetAllProducts(t *testing.T) {
	setup()
	expectedProducts := []models.Product{
		{ID: 1, Description: "Product 1", ProductType: "Type 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Description: "Product 2", ProductType: "Type 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	mockRepo.On("GetAllProducts", db).Return(expectedProducts, nil)

	products, err := GetAllProducts()
	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
	mockRepo.AssertExpectations(t)
}

func TestGetProductByDescription(t *testing.T) {
	setup()
	expectedProduct := models.Product{ID: 1, Description: "Product 1", ProductType: "Type 1", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	mockRepo.On("GetProductByDescription", db, "Product 1").Return(expectedProduct, nil)

	product, err := GetProductByDescription("Product 1")
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
	mockRepo.AssertExpectations(t)
}

func TestCreateProduct(t *testing.T) {
	setup()
	newProduct := models.Product{Description: "New Product", ProductType: "New Type"}
	mockRepo.On("CreateProduct", db, newProduct).Return(nil)

	err := CreateProduct(newProduct)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	setup()
	updatedProduct := models.Product{ID: 1, Description: "Updated Product", ProductType: "Updated Type"}
	mockRepo.On("UpdateProduct", db, updatedProduct).Return(nil)

	err := UpdateProduct(updatedProduct)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	setup()
	mockResult := sqlmock.NewResult(1, 1)
	mockRepo.On("DeleteProduct", db, 1).Return(mockResult, nil)

	result, err := DeleteProduct(1)
	assert.NoError(t, err)
	rowsAffected, err := result.RowsAffected()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowsAffected)
	mockRepo.AssertExpectations(t)
}
