package services

import (
	"API/internal/models"
	"API/internal/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.productRepo.GetProducts()
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.productRepo.CreateProduct(product)
}
