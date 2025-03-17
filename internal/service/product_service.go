package service

import (
	"github.com/ruanv123/api-go-crud/internal/model"
	"github.com/ruanv123/api-go-crud/internal/repository"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetAllProducts(name string, page, limit int) ([]model.Product, int64, error)
	GetProductByID(id uint) (*model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id uint) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo}
}

func (s *productService) CreateProduct(product *model.Product) error {
	return s.productRepo.Create(product)
}

// GetAllProducts retorna todos os produtos
func (s *productService) GetAllProducts(name string, page, limit int) ([]model.Product, int64, error) {
	return s.productRepo.FindAll(name, page, limit)
}

// GetProductByID retorna um produto pelo ID
func (s *productService) GetProductByID(id uint) (*model.Product, error) {
	return s.productRepo.FindByID(id)
}

// UpdateProduct atualiza um produto existente
func (s *productService) UpdateProduct(product *model.Product) error {
	return s.productRepo.Update(product)
}

// DeleteProduct exclui um produto pelo ID
func (s *productService) DeleteProduct(id uint) error {
	return s.productRepo.Delete(id)
}
