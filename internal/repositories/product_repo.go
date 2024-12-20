package repositories

import (
	"API/internal/models"
	"API/pkg/database"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := database.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) CreateProduct(product *models.Product) error {
	return database.DB.Create(product).Error
}
